package daos

import (
	"strings"

	"github.com/constant-money/constant-event/models"
	"github.com/jinzhu/gorm"
)

type MasterAddressDAO struct{}

func (ma *MasterAddressDAO) Update(masterAddress *models.MasterAddress, tx *gorm.DB) error {
	err := tx.Save(masterAddress).Error
	return err
}

func (ma *MasterAddressDAO) New(masterAddress *models.MasterAddress, tx *gorm.DB) error {
	err := tx.Create(masterAddress).Error
	return err
}

func (ma *MasterAddressDAO) GetAdddressReady() (models.MasterAddress, error) {
	masterAddress := models.MasterAddress{}
	err := models.Database().
		Where(`status = ?`, models.MasterAddressStatusReady).
		First(&masterAddress).Error

	if err != nil {
		return masterAddress, err
	}
	return masterAddress, nil
}

func (ma *MasterAddressDAO) UpdateStatusByTnxHash(tnxHash string, status models.MasterAddressStatus, tx *gorm.DB) (err error) {
	var address models.MasterAddress
	err = tx.Model(&address).Where("last_tnx_hash = ?", strings.ToLower(tnxHash)).Update("status", status).Error
	return
}

func (t *MasterAddressDAO) DeleteAll(query string, tx *gorm.DB) error {
	// This comment only update delete_at
	// err := tx.Where(query).Delete(models.MasterAddressDAO{}).Error
	err := tx.Unscoped().Where(query).Delete(models.MasterAddress{}).Error
	return err
}
