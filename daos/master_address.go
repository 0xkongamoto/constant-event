package daos

import (
	"github.com/constant-money/constant-event/models"
	wm "github.com/constant-money/constant-web-api/models"
	"github.com/jinzhu/gorm"
)

type MasterAddressDAO struct{}

func (ma *MasterAddressDAO) Update(masterAddress *wm.MasterAddress, tx *gorm.DB) error {
	err := tx.Save(masterAddress).Error
	return err
}

// GetAdddressReady : ...
func (ma *MasterAddressDAO) GetAdddressReady() ([]*wm.MasterAddress, error) {
	masterAddress := []*wm.MasterAddress{}
	err := models.Database().
		Where(`status = ?`, wm.MasterAddressStatusReady).
		Find(&masterAddress).Error

	if err != nil {
		return nil, err
	}
	return masterAddress, nil
}

// GetMasterAddress : ...
func (ma *MasterAddressDAO) GetMasterAddress() (*wm.MasterAddress, error) {
	masterAddress := wm.MasterAddress{}
	err := models.Database().First(&masterAddress).Error
	if err != nil {
		return nil, err
	}
	return &masterAddress, nil
}
