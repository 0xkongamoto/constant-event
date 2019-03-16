package ethereum

import (
	"fmt"
	"log"
	"math/big"

	contract "github.com/constant-money/constant-event/ethereum/contract"

	helpers "github.com/constant-money/constant-web-api/helpers"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Constant : struct
type Constant struct {
	ContractAddress         string
	ContractOwnerPrivateKey string
	ethereumService         *Ethereum
}

// InitConstant : contractAddr, ownerPriKey, ethereum
func InitConstant(contractAddr string, ownerPriKey string, cipherKey string, ethereum *Ethereum) *Constant {
	if cipherKey != "" {
		priKey, err := helpers.DecryptToString(ownerPriKey, cipherKey)
		if err != nil {
			priKey = ownerPriKey
		}

		return &Constant{
			ContractAddress:         contractAddr,
			ContractOwnerPrivateKey: priKey,
			ethereumService:         ethereum,
		}
	}

	return &Constant{
		ContractAddress:         contractAddr,
		ContractOwnerPrivateKey: ownerPriKey,
		ethereumService:         ethereum,
	}
}

// GetInstance : Constant
func (c *Constant) GetInstance() (*contract.Constant, error) {
	if c.isValid() {
		address := common.HexToAddress(c.ContractAddress)
		client, _ := c.ethereumService.GetClient()
		instance, err := contract.NewConstant(address, client)
		if err != nil {
			return nil, err
		}

		return instance, nil
	}

	return nil, nil
}

func (c *Constant) isValid() bool {
	if c.ContractAddress == "" ||
		c.ContractOwnerPrivateKey == "" {
		return false
	}

	return true
}

// Purchase : address, value, offchain
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
	fmt.Println(" DEBUG ", fromAddr, toAddr, value, offchain)
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

	fmt.Println("DEBUG 88")
	tx, err := instance.TransferByAdmin(auth, common.HexToAddress(fromAddr), common.HexToAddress(toAddr), value, o)
	fmt.Println("DEBUG 99")

	if err != nil {
		fmt.Println("DEBUG 100 ", err)
		log.Fatal(err)
		return "", err
	}
	return tx.Hash().Hex(), nil
}

// BalanceOf : address
func (c *Constant) BalanceOf(address string) (*big.Int, error) {
	instance, err := c.GetInstance()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	bal, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return bal, nil
}
