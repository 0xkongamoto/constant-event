package ethereum

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"log"
	"math/big"

	"github.com/constant-money/constant-event/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

const GasPriceDefault = 30000000000

type Ethereum struct {
	EthChainEnpoint string
	SimpleLoanOwner string
	SimpleLoadAddr  string
	client          *ethclient.Client
}

func Init(conf *config.Config) *Ethereum {
	service := &Ethereum{
		EthChainEnpoint: conf.ChainURL,
		// SimpleLoanOwner: conf.SimpleLoanOwner,
		// SimpleLoadAddr:  conf.SimpleLoanAddr,
	}
	return service
}

func (s *Ethereum) GetClient() (*ethclient.Client, error) {
	if s.client != nil {
		return s.client, nil
	}

	client, err := ethclient.Dial(s.EthChainEnpoint)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (s *Ethereum) GetGasPrice() (*big.Int, error) {
	client, err := s.GetClient()
	if err != nil {
		return big.NewInt(GasPriceDefault), err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return big.NewInt(GasPriceDefault), err
	}

	return gasPrice, nil
}

func (s *Ethereum) CreateWallet() (string, string, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	privateKey := hex.EncodeToString(key.D.Bytes())

	return address, privateKey, nil
}

// GetTransactionOpt : private key, address, value, data
func (s *Ethereum) GetTransactionOpt(fromPrvKey string) (*bind.TransactOpts, error) {
	client, err := s.GetClient()

	privateKey, err := crypto.HexToECDSA(fromPrvKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	gasPrice, err := s.GetGasPrice()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	return auth, nil
}

// SendSignedTransaction : private key, address, value, data
func (s *Ethereum) SendSignedTransaction(fromPrvKey string, to string, value *big.Int, data []byte) (string, error) {
	client, err := s.GetClient()

	if err != nil {
		return "", err
	}

	privateKey, err := crypto.HexToECDSA(fromPrvKey)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasLimit := uint64(300000) // in units
	gasPrice, err := s.GetGasPrice()
	if err != nil {
		return "", err
	}

	toAddress := common.HexToAddress(to)
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	ts := &(types.Transactions{signedTx})
	rawTxBytes := ts.GetRlp(0)

	rtx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &rtx)
	err = client.SendTransaction(context.Background(), rtx)
	if err != nil {
		return "", err
	}
	return rtx.Hash().Hex(), nil
}
