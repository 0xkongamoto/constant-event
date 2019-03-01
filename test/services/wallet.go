package services_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	_ "github.com/go-sql-driver/mysql"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type WalletTestSuite struct {
	suite.Suite
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *WalletTestSuite) SetupTest() {

}

func (suite *WalletTestSuite) TestWalletSrvInitSuccessfully() {

}

func TestWalletTestSuite(t *testing.T) {
	suite.Run(t, new(WalletTestSuite))
}
