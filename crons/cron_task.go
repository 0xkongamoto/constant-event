package crons

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/constant-money/constant-event/config"
	"github.com/constant-money/constant-event/daos"
	"github.com/constant-money/constant-event/ethereum"
	"github.com/constant-money/constant-event/models"
	wm "github.com/constant-money/constant-web-api/models"
	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
)

// CronTask : struct
type CronTask struct {
	ScanRunning      bool
	masterAddressDAO *daos.MasterAddressDAO
	taskDAO          *daos.TaskDAO
	txDAO            *daos.TxDAO
	conf             *config.Config
	etherSrv         *ethereum.Ethereum
	lastIdx          uint
}

// NewCronTask : addressDAO, taskDAO, txDAO, etherSrv, config
func NewCronTask(masterAddressDAO *daos.MasterAddressDAO, taskDao *daos.TaskDAO, txDAO *daos.TxDAO, etherSrv *ethereum.Ethereum, conf *config.Config) (crt *CronTask) {
	crt = &CronTask{
		masterAddressDAO: masterAddressDAO,
		taskDAO:          taskDao,
		txDAO:            txDAO,
		etherSrv:         etherSrv,
		conf:             conf,
	}
	return crt
}

// ScanTask : ...
func (cr *CronTask) ScanTask() {
	arrAddr, errMasterAddr := cr.masterAddressDAO.GetAdddressReady()
	if errMasterAddr != nil {
		log.Println("Get master address ready error: ", errMasterAddr.Error())
		return
	}

	tasks, errTasks := cr.taskDAO.GetTasksScanning(cr.lastIdx, len(arrAddr))
	if errTasks != nil {
		log.Println("Get Tasks error", errTasks.Error())
		return
	}

	if len(tasks) == 0 {
		fmt.Println("Tasks not found!!!")
		return
	}

	if len(tasks) > len(arrAddr) {
		log.Println("len(tasks) > len(arrAddr)")
		return
	}

	for i := 0; i < len(tasks); i++ {
		address := arrAddr[i]
		task := tasks[i]

		fmt.Println("DEBUG A")
		dataBytes := []byte(task.Data)
		var dataJSON map[string]interface{}
		if errUnmarshal := json.Unmarshal(dataBytes, &dataJSON); errUnmarshal != nil {
			log.Println("Unmarshal task data", errUnmarshal.Error())
			return
		}
		fmt.Println("DEBUG B")

		errOnchain := cr.handleSmartContractMethod(dataJSON, &task, address, cr.etherSrv, task.Method)

		if errOnchain == nil {
			cr.updateMasterAddrStatus(address, wm.MasterAddressStatusProgressing)
		}
		fmt.Println("DEBUG C")

		cr.lastIdx = task.ID

	}

}

func (cr *CronTask) handleSmartContractMethod(dataJSON map[string]interface{}, task *wm.Task, masterAddrReady *wm.MasterAddress, etherService *ethereum.Ethereum, method wm.TaskMethod) error {

	fmt.Println("DEBUG 11")
	dataJSON["ContractAddress"] = task.ContractAddress
	dataJSON["ContractName"] = task.ContractName
	dataJSON["MasterAddr"] = masterAddrReady.Address

	// TODO: select InitContract's version by name
	constantService := ethereum.InitConstant(task.ContractAddress, masterAddrReady.PriKey, cr.conf.CipherKey, etherService)

	var tnxHash string
	var errOnchain error

	fmt.Println("DEBUG 22")
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

	cr.updateTask(task, taskStatus)
	cr.saveTnx(tnxHash, string(dataStr), -1, dataJSON["Offchain"].(string), dataJSON["ContractAddress"].(string), string(task.Method), masterAddrReady.Address, task.ID)
	return errOnchain
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

func (cr *CronTask) updateTask(task *wm.Task, status wm.TaskStatus) error {
	errTx := models.WithTransaction(func(tx *gorm.DB) error {
		task.Status = status

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

func (cr *CronTask) updateMasterAddrStatus(masterAddr *wm.MasterAddress, status wm.MasterAddressStatus) error {
	errTx := models.WithTransaction(func(tx *gorm.DB) error {
		masterAddr.Status = status

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
	newTx := &wm.Tx{
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
