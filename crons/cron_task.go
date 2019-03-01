package crons

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/constant-money/constant-event/config"
	"github.com/constant-money/constant-event/daos"
	"github.com/constant-money/constant-event/ethereum"
	"github.com/constant-money/constant-event/models"
	wm "github.com/constant-money/constant-web-api/models"
	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
)

type CronTask struct {
	ScanRunning      bool
	QueueRange       int
	masterAddressDAO *daos.MasterAddressDAO
	taskDAO          *daos.TaskDAO
	txDAO            *daos.TxDAO
	conf             *config.Config
}

func NewCronTask(QueueRange int, masterAddressDAO *daos.MasterAddressDAO, taskDao *daos.TaskDAO, txDAO *daos.TxDAO, conf *config.Config) (crt CronTask) {
	crt = CronTask{false, QueueRange, masterAddressDAO, taskDao, txDAO, config.GetConfig()}
	return crt
}

func (cr *CronTask) ScanTask() {
	masterAddrReady, errMasterAddr := cr.masterAddressDAO.GetAdddressReady()
	if errMasterAddr != nil {
		log.Println("Get master address ready error", errMasterAddr.Error())
		return
	}

	flagAddrExist := false
	priKey := ""
	for _, masterAddr := range cr.conf.MasterAddresses {
		if strings.ToLower(masterAddrReady.Address) == strings.ToLower(masterAddr.Address) {
			flagAddrExist = true
			priKey = masterAddr.Key
			break
		}
	}

	if !flagAddrExist {
		log.Println("Master address not found in config")
		return
	}

	tasksBegin, errBegin := cr.taskDAO.GetLastIdScanning()
	var idBegin = tasksBegin.ID

	if errBegin != nil {
		log.Println("Get last Task error", errBegin.Error())
	}

	tasks, errTasks := cr.taskDAO.GetTasksScanning(idBegin, cr.QueueRange)
	if errTasks != nil {
		log.Println("Get Tasks error", errTasks.Error())
		return
	}

	if len(tasks) == 0 {
		fmt.Println("Tasks not found!!!")
		return
	}

	etherService := ethereum.Init(cr.conf)

	for _, task := range tasks {
		// TODO: check task's MasterAddr is ready
		errUpdate := cr.updateTask(&task, masterAddrReady.Address, wm.TaskStatusProgressing)
		if errUpdate != nil {
			continue
		}

		dataBytes := []byte(task.Data)
		var dataJSON map[string]interface{}
		if errUnmarshal := json.Unmarshal(dataBytes, &dataJSON); errUnmarshal != nil {
			log.Println("Unmarshal task data", errUnmarshal.Error())
			return
		}

		var errOnchain error
		var tnxHash string

		tnxHash, errOnchain = cr.handleSmartContractMethod(dataJSON, &task, masterAddrReady, priKey, etherService, task.Method)

		if errOnchain == nil {
			cr.updateMasterAddrStatus(masterAddrReady, wm.MasterAddressStatusProgressing, tnxHash)
		}
	}
}

func (cr *CronTask) handleSmartContractMethod(dataJSON map[string]interface{}, task *wm.Task, masterAddrReady *wm.MasterAddress, priKey string, etherService *ethereum.Ethereum, method wm.TaskMethod) (string, error) {
	dataJSON["ContractAddress"] = task.ContractAddress
	dataJSON["ContractName"] = task.ContractName
	dataJSON["MasterAddr"] = masterAddrReady.Address

	// TODO: select InitContract's version by name
	constantService := ethereum.InitConstant(task.ContractAddress, priKey, etherService)

	var tnxHash string
	var errOnchain error

	switch task.Method {

	case wm.TaskMethodPurchase:
		var data models.PurchaseParams
		mapstructure.Decode(dataJSON, &data)
		tnxHash, errOnchain = cr.handlePurchase(&data, task.ID, constantService)

	case wm.TaskMethodRedeem:
		var data models.RedeemParams
		mapstructure.Decode(dataJSON, &data)
		tnxHash, errOnchain = cr.handleRedeem(&data, task.ID, constantService)

	case wm.TaskMethodTransferByAdmin:
		var data models.TransferByAdminParams
		mapstructure.Decode(dataJSON, &data)
		tnxHash, errOnchain = cr.handleTransferByAdmin(&data, task.ID, constantService)
	}

	var taskStatus wm.TaskStatus

	if taskStatus = wm.TaskStatusSuccess; errOnchain != nil {
		taskStatus = wm.TaskStatusFailed
		dataJSON["Err"] = errOnchain.Error()
	}
	dataStr, _ := json.Marshal(dataJSON)
	// task.Data = string(dataStr)

	cr.updateTask(task, masterAddrReady.Address, taskStatus)
	cr.saveTnx(tnxHash, string(dataStr), -1, dataJSON["Offchain"].(string), dataJSON["ContractAddress"].(string), string(task.Method), masterAddrReady.Address, task.ID)
	return tnxHash, errOnchain
}

func (cr *CronTask) handlePurchase(params *models.PurchaseParams, taskID uint, constantService *ethereum.Constant) (string, error) {
	value := new(big.Int)
	value, ok := value.SetString(params.Value, 10)
	if !ok {
		log.Println("Purchase SetString: error")
		return "", errors.New("Purchase SetString: error")
	}
	tnxHash, err := constantService.Purchase(params.Purchaser, value, params.Offchain)
	return tnxHash, err
}

func (cr *CronTask) handleRedeem(params *models.RedeemParams, taskID uint, constantService *ethereum.Constant) (string, error) {
	value := new(big.Int)
	value, ok := value.SetString(params.Value, 10)
	if !ok {
		log.Println("Redeem SetString: error")
		return "", errors.New("Purchase SetString: error")
	}
	tnxHash, err := constantService.Redeem(params.Redeemer, value, params.Offchain)
	return tnxHash, err
}

func (cr *CronTask) handleTransferByAdmin(params *models.TransferByAdminParams, taskID uint, constantService *ethereum.Constant) (string, error) {
	value := new(big.Int)
	value, ok := value.SetString(params.Value, 10)
	if !ok {
		log.Println("TransferByAdmin SetString: error")
		return "", errors.New("TransferByAdmin SetString: error")
	}
	tnxHash, err := constantService.TransferByAdmin(params.FromAddress, params.ToAddress, value, params.Offchain)
	return tnxHash, err
}

func (cr *CronTask) updateTask(task *wm.Task, masterAddr string, status wm.TaskStatus) error {
	errTx := models.WithTransaction(func(tx *gorm.DB) error {
		task.Status = status
		task.MasterAddress = masterAddr

		if err := cr.taskDAO.Update(task, tx); err != nil {
			log.Println("Update Task error", err.Error())
			return err
		}
		return nil
	})

	if errTx != nil {
		log.Println("DB Tnx Update Task error", errTx.Error())
	}
	return errTx
}

func (cr *CronTask) updateMasterAddrStatus(masterAddr *wm.MasterAddress, status wm.MasterAddressStatus, tnxHash string) error {
	errTx := models.WithTransaction(func(tx *gorm.DB) error {
		masterAddr.Status = status

		if tnxHash != "" {
			masterAddr.LastTnxHash = strings.ToLower(tnxHash)
		}

		if err := cr.masterAddressDAO.Update(masterAddr, tx); err != nil {
			log.Println("Update Master Address Status error", err.Error())
			return err
		}
		return nil
	})

	if errTx != nil {
		log.Println("DB Tnx Update Master Address Status error", errTx.Error())
	}
	return errTx
}

func (cr *CronTask) saveTnx(hash string, payload string, status int, offchain string, constractAddr string, contractMethod string, masterAddress string, taskID uint) error {
	newTx := &models.Tx{
		Hash:            hash,
		Payload:         payload,
		Status:          status,
		ChainID:         cr.conf.ChainID,
		Offchain:        offchain,
		TaskID:          taskID,
		MasterAddress:   masterAddress,
		ContractAddress: constractAddr,
		ContractMethod:  contractMethod,
	}
	err := cr.txDAO.New(newTx)
	if err != nil {
		log.Println("DB Tnx Save Tx error", err.Error())
	}
	return err
}
