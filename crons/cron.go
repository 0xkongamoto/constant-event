package crons

import (
	"context"
	"encoding/json"
	"log"

	"github.com/constant-money/constant-event/config"
	"github.com/constant-money/constant-event/daos"
	"github.com/constant-money/constant-event/models"
	"github.com/constant-money/constant-event/services"
	"github.com/constant-money/constant-event/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"
)

var txDAO = &daos.TxDAO{}
var hookService = &services.HookService{}
var etherscanService = &services.EtherscanService{}

// Cron : main struct to handle scan and sync
type Cron struct {
	ScanRunning      bool
	SyncRunning      bool
	ContractJSON     string
	ContractAddress  string
	masterAddressDAO *daos.MasterAddressDAO
}

// NewCron : creates a new Cron instance
func NewCron(contractJSON string, contractAddress string, masterAddressDAO *daos.MasterAddressDAO) (cr Cron) {
	cr = Cron{false, false, contractJSON, contractAddress, masterAddressDAO}
	return
}

// ScanTx : load pending tx from db
func (cr *Cron) ScanTx() {
	// todo get list transaction pending
	query := "hash != -1 and status = -1 and contract_address like '" + cr.ContractAddress + "'"
	transactions, err := txDAO.GetTxPending(query)
	if err != nil {
		log.Println("Scan Tx error", err.Error())
		return
	}
	if len(transactions) == 0 {
		log.Printf("Scan Tx: don't have any pending tx for contract %s\n", cr.ContractAddress)
		return
	}

	log.Printf("Have %d pending tx\n", len(transactions))

	conf := config.GetConfig()
	networkURL := conf.ChainURL

	etherClient, err := ethclient.Dial(networkURL)
	if err != nil {
		log.Printf("Scan Tx: connect to network %s fail!\n", networkURL)
		return
	}

	totalJobs := len(transactions)
	jobs := make(chan models.Tx, 100)
	results := make(chan bool, totalJobs)

	workers := totalJobs / 10
	if workers > 50 {
		workers = 50
	}
	if workers == 0 {
		workers = 1
	}

	for w := 1; w <= workers; w++ {
		go cr.scanWorker(w, etherClient, jobs, results)
	}
	// todo loop & parse transaction
	for _, transaction := range transactions {
		jobs <- transaction
	}
	close(jobs)

	for i := 0; i < totalJobs; i++ {
		<-results
	}
	log.Println("scan complete")
}

// SyncTx : call etherscan to make sure not miss any tx
func (cr *Cron) SyncTx() {
	// todo call etherscan.io to get all transactions
	conf := config.GetConfig()
	chainID := conf.ChainID

	page := 1
	recordPerPage := 100
	for {
		status, transactions := etherscanService.ListTransactions(cr.ContractAddress, page, recordPerPage)
		log.Println(status, len(transactions), page, recordPerPage)
		if !status {
			log.Println("etherscan.io return error")
			break
		}
		if len(transactions) == 0 {
			break
		}

		page = page + 1
		for _, transaction := range transactions {
			transactionObj := transaction.(map[string]interface{})
			hash := transactionObj["hash"].(string)
			input := transactionObj["input"].(string)

			_, err := txDAO.GetByHash(hash)
			if err != nil {
				status, inputJSON := utils.DecodeTransactionInput(cr.ContractJSON, input)
				if status {
					var jsonData map[string]interface{}
					json.Unmarshal([]byte(inputJSON), &jsonData)

					offchain, hasOffchain := jsonData["offchain"]

					tx := &models.Tx{
						Hash:            hash,
						ContractAddress: cr.ContractAddress,
						ContractMethod:  jsonData["methodName"].(string),
						Payload:         input,
						ChainID:         chainID,
					}

					if hasOffchain {
						tx.Offchain = offchain.(string)
					}
					err := txDAO.New(tx)
					if err != nil {
						log.Println("Sync new transaction error", err.Error())
					}
				}
			}
		}
	}
}

func (cr *Cron) scanWorker(id int, etherClient *ethclient.Client, jobs <-chan models.Tx, results chan<- bool) {
	for transaction := range jobs {
		log.Printf("start scan %s\n", transaction.Hash)
		txHash := common.HexToHash(transaction.Hash)
		tx, pending, err := etherClient.TransactionByHash(context.Background(), txHash)
		if err == nil && !pending {
			receipt, err := etherClient.TransactionReceipt(context.Background(), txHash)
			if err != nil {
				log.Println("Scan Tx: get receipt error", err.Error())
			} else {
				log.Printf("Tx %s has receipt, status %d\n", transaction.Hash, receipt.Status)

				cr.updateMasterAddrStatus(transaction.Hash, models.MasterAddressStatusReady)

				if receipt.Status == 0 {
					// case fail
					decodeStatus, methodJSON := utils.DecodeTransactionInput(cr.ContractJSON, common.ToHex(tx.Data()))
					// call REST fail
					var jsonData map[string]interface{}
					json.Unmarshal([]byte(methodJSON), &jsonData)
					jsonData["ID"] = transaction.ID
					status := 0
					if !decodeStatus {
						status = 2
					}
					jsonData["Status"] = status
					jsonData["ContractName"] = cr.ContractJSON
					jsonData["ContractAddress"] = cr.ContractAddress
					jsonData["Hash"] = transaction.Hash

					jsonWebhook := make(map[string]interface{})
					jsonWebhook["type"] = 0 /* WebhookTypeTxHash */
					jsonWebhook["data"] = jsonData

					err := hookService.Event(jsonWebhook)
					if err != nil {
						log.Println("Hook event fail error: ", err.Error())
						log.Println(jsonWebhook)
					}

				} else if receipt.Status == 1 {
					// case success
					log.Printf("Tx %s has receipt, logs %d\n", transaction.Hash, len(receipt.Logs))

					cr.updateMasterAddrStatus(transaction.Hash, models.MasterAddressStatusReady)

					if len(receipt.Logs) > 0 {
						for _, l := range receipt.Logs {
							decodeStatus, eventJSON := utils.DecodeTransactionLog(cr.ContractJSON, l)
							if eventJSON != "" {
								var jsonData map[string]interface{}
								json.Unmarshal([]byte(eventJSON), &jsonData)
								jsonData["ID"] = transaction.ID
								jsonData["Status"] = 1
								jsonData["ContractName"] = cr.ContractJSON
								jsonData["ContractAddress"] = cr.ContractAddress
								jsonData["Hash"] = transaction.Hash

								jsonWebhook := make(map[string]interface{})
								jsonWebhook["type"] = 0 /* WebhookTypeTxHash */
								jsonWebhook["data"] = jsonData

								if decodeStatus {
									// call REST API SUCCESS with event
									//log.Println("hook success", jsonData)
									err := hookService.Event(jsonWebhook)
									if err != nil {
										log.Println("Hook event success error: ", err.Error())
									}
								}
								log.Println(jsonWebhook)
							}
						}
					}
				} else {
					log.Println("Unknown case")
				}
			}
		} else {
			log.Printf("Tx %s is pending or error occured\n", transaction.Hash)
		}
		results <- true
	}
}

func (cr *Cron) updateMasterAddrStatus(tnxHash string, status models.MasterAddressStatus) error {
	errTx := models.WithTransaction(func(tx *gorm.DB) error {
		if err := cr.masterAddressDAO.UpdateStatusByTnxHash(tnxHash, models.MasterAddressStatusReady, tx); err != nil {
			log.Println("Update Master Address Ready error", err.Error())
			return err
		}
		return nil
	})

	if errTx != nil {
		log.Println("DB Tnx Update Master Address Ready error", errTx.Error())
	}
	return errTx
}