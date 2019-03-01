package daos

import (
	"strings"

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
func (ma *MasterAddressDAO) GetAdddressReady() (*wm.MasterAddress, error) {
	masterAddress := wm.MasterAddress{}
	err := models.Database().
		Where(`status = ?`, wm.MasterAddressStatusReady).
		First(&masterAddress).Error

	if err != nil {
		return nil, err
	}
	return &masterAddress, nil
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

func (ma *MasterAddressDAO) UpdateStatusByTnxHash(tnxHash string, status wm.MasterAddressStatus, tx *gorm.DB) (err error) {
	var address wm.MasterAddress
	err = tx.Model(&address).Where("last_tnx_hash = ?", strings.ToLower(tnxHash)).Update("status", status).Error
	return
}

func (t *MasterAddressDAO) DeleteAll(query string, tx *gorm.DB) error {
	err := tx.Unscoped().Where(query).Delete(wm.MasterAddress{}).Error
	return err
}
