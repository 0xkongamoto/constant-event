package config

import (
	"encoding/json"
	"log"
	"os"
)

var config *Config

func init() {
	file, err := os.Open("config/conf.json")
	if err != nil {
		log.Println("error:", err)
		panic(err)
	}
	decoder := json.NewDecoder(file)
	v := Config{}
	err = decoder.Decode(&v)
	if err != nil {
		log.Println("error:", err)
		panic(err)
	}
	config = &v
}

func GetConfig() *Config {
	return config
}

type Config struct {
	Port                int        `json:"port"`
	Db                  string     `json:"db"`
	Env                 string     `json:"env"`
	ChainID             uint       `json:"chain_id"`
	ChainURL            string     `json:"chain_url"`
	EtherscanURL        string     `json:"etherscan_url"`
	EtherscanKey        string     `json:"etherscan_key"`
	HookEndpoint        string     `json:"hook_endpoint"`
	PrimetrustEndpoint  string     `json:"primetrust_prefix"`
	PrimetrustUsername  string     `json:"primetrust_email"`
	PrimetrustPassword  string     `json:"primetrust_password"`
	PrimetrustAccountID string     `json:"primetrust_account_id"`
	Contracts           []Contract `json:"contracts"`
	CipherKey           string     `json:"wallet_cipher_key"`
	BlockexplorerAPI    string     `json:"blockexp_api_url"`
	BlockCypherAPI      string     `json:"blockcyp_api_url"`
	BtcIsTestnet        bool       `json:"btc_is_testnet"`
	BcyToken            string     `json:"bcy_token"`
	MasterEthWallet     string     `json:"master_eth_wallet"`
	MasterBtcWallet     string     `json:"master_btc_wallet"`
	MasterUsdtWallet    string     `json:"master_usdt_wallet"`
}

type Contract struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
