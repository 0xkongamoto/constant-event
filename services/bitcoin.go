package services

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/constant-money/constant-event/config"
	helpers "github.com/constant-money/constant-web-api/helpers"
)

type Transaction struct {
	TxID               string `json:"txid"`
	SourceAddress      string `json:"source_address"`
	DestinationAddress string `json:"destination_address"`
	Amount             int64  `json:"amount"`
	UnsignedTx         string `json:"unsignedtx"`
	SignedTx           string `json:"signedtx"`
}

type UTXO struct {
	Address       string  `json:"address"`
	TxID          string  `json:"txid"`
	Vout          uint    `json:"vout"`
	ScriptPubKey  string  `json:"scriptPubKey"`
	Amount        float64 `json:"amount"`
	Satoshis      uint64  `json:"satoshis"`
	Height        uint64  `json:"height"`
	Confirmations int     `json:"confirmations"`
}

type Txref struct {
	TxHash        string  `json:"tx_hash"`
	BlockHeight   string  `json:"block_height"`
	TxInputN      uint    `json:"tx_input_n"`
	txOutputN     string  `json:"tx_output_n"`
	Value         float64 `json:"value"`
	RefBalance    uint64  `json:"ref_balance"`
	Spent         uint64  `json:"spent"`
	Confirmations int     `json:"confirmations"`
	Confirmed     int     `json:"confirmed"`
	DoubleSpend   int     `json:"double_spend"`
}

type AddrInfo struct {
	Address            string  `json:"address"`
	TotalReceived      uint64  `json:"total_received"`
	TotalSent          uint64  `json:"total_sent"`
	Balance            uint64  `json:"balance"`
	UnconfirmedBalance uint64  `json:"unconfirmed_balance"`
	FinalBalance       uint64  `json:"final_balance"`
	NTx                uint    `json:"n_tx"`
	UnconfirmedNTx     uint    `json:"unconfirmed_n_tx"`
	FinalNTX           uint    `json:"final_n_tx"`
	TxRefs             []Txref `json:"txrefs"`
}

// BitcoinService :
type BitcoinService struct {
	conf *config.Config
}

func NewBitcoinService(conf *config.Config) *BitcoinService {
	return &BitcoinService{
		conf: conf,
	}
}

// BTCBalanceOf :
func (bs *BitcoinService) BTCBalanceOf(address string) (string, error) {
	url := fmt.Sprintf("%s/addr/%s/balance", bs.conf.BlockexplorerAPI, address)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("BTC get balanceOf failed", address, err.Error())
		return "", err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Call BTC get balanceOf failed", err.Error())
		return "", err
	}
	defer func(r *http.Response) {
		err := r.Body.Close()
		if err != nil {
			fmt.Println("Close body failed", err.Error())
		}
	}(res)

	if res.StatusCode != http.StatusOK {
		fmt.Println("Balanceof Response status != 200")
		return "", errors.New("Response status != 200")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Read body failed", err.Error())
		return "", errors.New("Read body failed")
	}
	return string(body), nil
}

// BTCGetUTXO :
func (bs *BitcoinService) BTCGetLastUTXO(address string) (string, error) {
	// url := fmt.Sprintf("%s/%s/full?limit=1&unspentOnly=true&includeScript=false", bs.conf.BlockCypherAPI, address)
	url := fmt.Sprintf("%s/%s?limit=1&unspentOnly=true&includeScript=false", bs.conf.BlockCypherAPI, address)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("BTC get UTXO failed", address, err.Error())
		return "", err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Call BTC get UTXO failed", err.Error())
		return "", err
	}

	defer func(r *http.Response) {
		err := r.Body.Close()
		if err != nil {
			fmt.Println("Close body failed", err.Error())
		}
	}(res)

	if res.StatusCode != http.StatusOK {
		return "", errors.New("GetUTXO Response status != 200")
	}

	body, err := ioutil.ReadAll(res.Body)

	var data AddrInfo
	json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Read body failed", err.Error())
		return "", errors.New("Read body failed")
	}
	if len(data.TxRefs) == 0 {
		return "", errors.New("GetUTXO TxRefs empty")
	}
	return data.TxRefs[0].TxHash, nil
}

// BTCSendRawTnx :
func (bs *BitcoinService) BTCSendRawTnx(from string, secret string, cipherKey string, destination string, amount int64) (string, error) {
	priKey, err := helpers.DecryptToString(secret, cipherKey)

	// Get balance of fromAddress
	balance, err := bs.BTCBalanceOf(from)
	if balance == "0" {
		return "", errors.New("Balance = 0")
	}

	// Get last txID
	utxoTxHash, err := bs.BTCGetLastUTXO(from)
	if err != nil {
		return "", err
	}

	// Generate transaction
	transacion, err := bs.createTransaction(priKey, destination, amount, utxoTxHash)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/tx/send", bs.conf.BlockexplorerAPI)
	jsonValue, _ := json.Marshal(map[string]interface{}{
		"rawtx": transacion.SignedTx,
	})
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("BTC send raw tnx failed", err.Error())
		return "", err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Call BTC send raw tnx failed", err.Error())
		return "", err
	}
	defer func(r *http.Response) {
		err := r.Body.Close()
		if err != nil {
			fmt.Println("Close body failed", err.Error())
		}
	}(res)

	if res.StatusCode != http.StatusOK {
		fmt.Println("Send RawTx Response status != 200")
		return "", errors.New("Send RawTx Response status != 200")
	}

	body, err := ioutil.ReadAll(res.Body)
	var data map[string]interface{}
	json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println("Read body failed", err.Error())
		return "", errors.New("Read body failed")
	}
	return string(body), nil
}

// CreateTransaction : ...
func (bs *BitcoinService) createTransaction(secret string, destination string, amount int64, txHash string) (Transaction, error) {
	var transaction Transaction
	var networkParams = &chaincfg.MainNetParams
	if bs.conf.BtcIsTestnet {
		networkParams = &chaincfg.TestNet3Params
	}

	wif, err := btcutil.DecodeWIF(secret)
	if err != nil {
		return Transaction{}, err
	}

	addresspubkey, _ := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeUncompressed(), networkParams)
	sourceTx := wire.NewMsgTx(wire.TxVersion)
	sourceUtxoHash, _ := chainhash.NewHashFromStr(txHash)
	sourceUtxo := wire.NewOutPoint(sourceUtxoHash, 0)
	sourceTxIn := wire.NewTxIn(sourceUtxo, nil, nil)
	destinationAddress, err := btcutil.DecodeAddress(destination, networkParams)
	sourceAddress, err := btcutil.DecodeAddress(addresspubkey.EncodeAddress(), networkParams)
	if err != nil {
		return Transaction{}, err
	}

	destinationPkScript, _ := txscript.PayToAddrScript(destinationAddress)
	sourcePkScript, _ := txscript.PayToAddrScript(sourceAddress)
	sourceTxOut := wire.NewTxOut(amount, sourcePkScript)
	sourceTx.AddTxIn(sourceTxIn)
	sourceTx.AddTxOut(sourceTxOut)
	sourceTxHash := sourceTx.TxHash()
	redeemTx := wire.NewMsgTx(wire.TxVersion)
	prevOut := wire.NewOutPoint(&sourceTxHash, 0)
	redeemTxIn := wire.NewTxIn(prevOut, nil, nil)
	redeemTx.AddTxIn(redeemTxIn)
	redeemTxOut := wire.NewTxOut(amount, destinationPkScript)
	redeemTx.AddTxOut(redeemTxOut)
	sigScript, err := txscript.SignatureScript(redeemTx, 0, sourceTx.TxOut[0].PkScript, txscript.SigHashAll, wif.PrivKey, false)
	if err != nil {
		return Transaction{}, err
	}

	redeemTx.TxIn[0].SignatureScript = sigScript
	flags := txscript.StandardVerifyFlags
	vm, err := txscript.NewEngine(sourceTx.TxOut[0].PkScript, redeemTx, 0, flags, nil, nil, amount)
	if err != nil {
		return Transaction{}, err
	}
	if err := vm.Execute(); err != nil {
		return Transaction{}, err
	}

	var unsignedTx bytes.Buffer
	var signedTx bytes.Buffer
	sourceTx.Serialize(&unsignedTx)
	redeemTx.Serialize(&signedTx)
	transaction.TxID = sourceTxHash.String()
	transaction.UnsignedTx = hex.EncodeToString(unsignedTx.Bytes())
	transaction.Amount = amount
	transaction.SignedTx = hex.EncodeToString(signedTx.Bytes())
	transaction.SourceAddress = sourceAddress.EncodeAddress()
	transaction.DestinationAddress = destinationAddress.EncodeAddress()
	return transaction, nil
}
