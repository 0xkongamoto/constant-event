package services

import (
	"fmt"

	"github.com/blockcypher/gobcy"
	"github.com/constant-money/constant-event/config"
	helpers "github.com/constant-money/constant-web-api/helpers"
)

// BlockcypherService : ...
type BlockcypherService struct {
	conf  *config.Config
	chain gobcy.API
}

// NewBlockcypherService : ...
func NewBlockcypherService(conf *config.Config) *BlockcypherService {
	// For Bitcoin main:
	chain := gobcy.API{conf.BcyToken, "btc", "main"}
	if conf.BtcIsTestnet {
		// For BlockCypher's internal testnet:
		chain = gobcy.API{conf.BcyToken, "bcy", "test"}
	}

	return &BlockcypherService{
		conf:  conf,
		chain: chain,
	}
}

// SendTX : ...
func (bs *BlockcypherService) SendTX(from string, secret string, cipherKey string, destination string, amount int) (string, error) {
	priKey, err := helpers.DecryptToString(secret, cipherKey)

	// addr1, _ := bs.chain.GenAddrKeychain()
	// addr2, _ := bs.chain.GenAddrKeychain()
	// _, _ = bs.chain.Faucet(addr1, 3e5)

	//Post New TXSkeleton
	skel, err := bs.chain.NewTX(gobcy.TempNewTX(from, destination, 2e5), false)
	//Sign it locally
	err = skel.Sign([]string{priKey})
	if err != nil {
		fmt.Println(err)
	}
	//Send TXSkeleton
	skel, err = bs.chain.SendTX(skel)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", skel)
	return "", nil
}
