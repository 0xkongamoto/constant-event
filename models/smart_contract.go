package models

type PurchaseParams struct {
	Purchaser       string `json:"Purchaser"`
	Value           string `json:"Value"`
	Offchain        string `json:"Offchain"`
	ContractAddress string `json:"ContractAddress"`
	ContractName    string `json:"ContractName"`
	MasterAddr      string `json:"MasterAddr"`
	Err             string
}

type RedeemParams struct {
	Redeemer        string `json:"Redeemer"`
	Value           string `json:"Value"`
	Offchain        string `json:"Offchain"`
	ContractAddress string `json:"ContractAddress"`
	ContractName    string `json:"ContractName"`
	MasterAddr      string `json:"MasterAddr"`
	Err             string
}

type TransferByAdminParams struct {
	FromAddress     string `json:"FromAddress"`
	ToAddress       string `json:"ToAddress"`
	Value           string `json:"Value"`
	Offchain        string `json:"Offchain"`
	ContractAddress string `json:"ContractAddress"`
	ContractName    string `json:"ContractName"`
	MasterAddr      string `json:"MasterAddr"`
	Err             string
}
