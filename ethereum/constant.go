package ethereum

import (
	"log"
	"math/big"

	contract "github.com/constant-money/constant-event/ethereum/contract"
	"github.com/ethereum/go-ethereum/common"
)

// Constant : struct
type Constant struct {
	ContractAddress         string
	ContractOwnerPrivateKey string
	ethereumService         *Ethereum
}

// InitConstant :
func InitConstant(contractAddr string, ownerPriKey string, ethereum *Ethereum) *Constant {
	c := &Constant{
		ContractAddress:         contractAddr,
		ContractOwnerPrivateKey: ownerPriKey,
		ethereumService:         ethereum,
	}
	return c
}

// GetInstance : Constant
func (c *Constant) GetInstance() (*contract.Constant, error) {
	address := common.HexToAddress(c.ContractAddress)
	client, _ := c.ethereumService.GetClient()
	instance, err := contract.NewConstant(address, client)
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func (c *Constant) Purchase(address string, value *big.Int, offchain string) (string, error) {
	instance, err := c.GetInstance()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	auth, err := c.ethereumService.GetTransactionOpt(c.ContractOwnerPrivateKey)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	o := [32]byte{}
	copy(o[:], []byte(offchain))

	tx, err := instance.Purchase(auth, common.HexToAddress(address), value, o)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (c *Constant) Redeem(address string, value *big.Int, offchain string) (string, error) {
	instance, err := c.GetInstance()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	auth, err := c.ethereumService.GetTransactionOpt(c.ContractOwnerPrivateKey)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	o := [32]byte{}
	copy(o[:], []byte(offchain))

	tx, err := instance.Redeem(auth, common.HexToAddress(address), value, o)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (c *Constant) TransferByAdmin(fromAddr string, toAddr string, value *big.Int, offchain string) (string, error) {
	instance, err := c.GetInstance()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	auth, err := c.ethereumService.GetTransactionOpt(c.ContractOwnerPrivateKey)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	o := [32]byte{}
	copy(o[:], []byte(offchain))

	tx, err := instance.TransferByAdmin(auth, common.HexToAddress(fromAddr), common.HexToAddress(toAddr), value, o)

	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return tx.Hash().Hex(), nil
}
