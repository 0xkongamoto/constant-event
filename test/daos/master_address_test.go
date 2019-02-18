package daos_test

import (
	"fmt"
	"testing"

	"github.com/constant-money/constant-event/daos"
	"github.com/constant-money/constant-event/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"

	_ "github.com/go-sql-driver/mysql"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type MasterAddressDAOTestSuite struct {
	suite.Suite
	masterAddressDAO *daos.MasterAddressDAO
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (ma *MasterAddressDAOTestSuite) SetupTest() {
	// init daos
	if db := models.Database(); db == nil {
		panic("Db cannot connect")
	}
}

func (suite *MasterAddressDAOTestSuite) TestAddNewMasterAddress() {
	newMasterAdr := &models.MasterAddress{
		Address:         "0x12345",
		Status:          models.MasterAddressStatusReady,
		LastTnxHash:     "0x1234",
		LastTnxTime:     999999,
		LastBlockNumber: 100,
		Nonce:           0,
	}

	errTx := models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.masterAddressDAO.New(newMasterAdr, tx)
		return err
	})

	suite.Nil(errTx)

	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.masterAddressDAO.DeleteAll(fmt.Sprintf("id = %v", newMasterAdr.ID), tx)
		return err
	})

	suite.Nil(errTx)
}

func (suite *MasterAddressDAOTestSuite) TestAddNewMasterAddressSameAddress() {
	newMasterAdr1 := &models.MasterAddress{
		Address:         "0x12345",
		Status:          models.MasterAddressStatusReady,
		LastTnxHash:     "0x1234",
		LastTnxTime:     999999,
		LastBlockNumber: 100,
		Nonce:           0,
	}

	errTx := models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.masterAddressDAO.New(newMasterAdr1, tx)
		return err
	})

	newMasterAdr2 := &models.MasterAddress{
		Address:         "0x12345",
		Status:          models.MasterAddressStatusReady,
		LastTnxHash:     "0x1234",
		LastTnxTime:     999999,
		LastBlockNumber: 100,
		Nonce:           0,
	}

	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.masterAddressDAO.New(newMasterAdr2, tx)
		return err
	})

	suite.NotNil(errTx)

	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.masterAddressDAO.DeleteAll(fmt.Sprintf("id IN (%d, %d)", newMasterAdr1.ID, newMasterAdr2.ID), tx)
		return err
	})

	suite.Nil(errTx)
}

func (suite *MasterAddressDAOTestSuite) TestUpdateNewMasterAddress() {
	newMasterAdr := &models.MasterAddress{
		Address:         "0x12345",
		Status:          models.MasterAddressStatusReady,
		LastTnxHash:     "0x1234",
		LastTnxTime:     999999,
		LastBlockNumber: 100,
		Nonce:           0,
	}

	errTx := models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.masterAddressDAO.New(newMasterAdr, tx)
		return err
	})

	suite.Nil(errTx)

	newMasterAdr.Address = "0x6789"
	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.masterAddressDAO.Update(newMasterAdr, tx)
		return err
	})

	suite.Nil(errTx)

	masterAdr := models.MasterAddress{}
	models.Database().Table("master_addresses").Where(`id=?`, newMasterAdr.ID).First(&masterAdr)

	suite.Equal(masterAdr.Address, "0x6789")

	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.masterAddressDAO.DeleteAll(fmt.Sprintf("id IN (%d)", newMasterAdr.ID), tx)
		return err
	})

	suite.Nil(errTx)
}

func (suite *MasterAddressDAOTestSuite) TestUpdateNewMasterAddressSameAddress() {
	newMasterAdr1 := &models.MasterAddress{
		Address:         "0x12345",
		Status:          models.MasterAddressStatusReady,
		LastTnxHash:     "0x1234",
		LastTnxTime:     999999,
		LastBlockNumber: 100,
		Nonce:           0,
	}

	errTx := models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.masterAddressDAO.New(newMasterAdr1, tx)
		return err
	})

	suite.Nil(errTx)

	newMasterAdr2 := &models.MasterAddress{
		Address:         "0x56789",
		Status:          models.MasterAddressStatusReady,
		LastTnxHash:     "0x1234",
		LastTnxTime:     999999,
		LastBlockNumber: 100,
		Nonce:           0,
	}

	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.masterAddressDAO.New(newMasterAdr2, tx)
		return err
	})

	suite.Nil(errTx)

	newMasterAdr1.Address = newMasterAdr2.Address
	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.masterAddressDAO.Update(newMasterAdr1, tx)
		return err
	})

	suite.NotNil(errTx)

	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.masterAddressDAO.DeleteAll(fmt.Sprintf("id IN (%d, %d)", newMasterAdr1.ID, newMasterAdr2.ID), tx)
		return err
	})

	suite.Nil(errTx)
}

func (suite *MasterAddressDAOTestSuite) TestGetAddressReady() {
	newMasterAdrrReady := &models.MasterAddress{
		Address:         "0x12345",
		Status:          models.MasterAddressStatusReady,
		LastTnxHash:     "0x1234",
		LastTnxTime:     999999,
		LastBlockNumber: 100,
		Nonce:           0,
	}

	errTx := models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.masterAddressDAO.New(newMasterAdrrReady, tx)
		return err
	})

	suite.Nil(errTx)

	newMasterAdrNotReady := &models.MasterAddress{
		Address:         "0x56789",
		Status:          models.MasterAddressStatusProgressing,
		LastTnxHash:     "0x1234",
		LastTnxTime:     999999,
		LastBlockNumber: 100,
		Nonce:           0,
	}

	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.masterAddressDAO.New(newMasterAdrNotReady, tx)
		return err
	})

	suite.Nil(errTx)

	masterAdress, errFind := suite.masterAddressDAO.GetAdddressReady()
	suite.Nil(errFind)
	suite.Equal(masterAdress.ID, newMasterAdrrReady.ID)

	errTx = models.WithTransaction(func(tx *gorm.DB) error {
		err := suite.masterAddressDAO.DeleteAll(fmt.Sprintf("id IN (%d, %d)", newMasterAdrrReady.ID, newMasterAdrNotReady.ID), tx)
		return err
	})

	suite.Nil(errTx)
}
func TestMasterAddressDAOTestSuite(t *testing.T) {
	suite.Run(t, new(MasterAddressDAOTestSuite))
}
