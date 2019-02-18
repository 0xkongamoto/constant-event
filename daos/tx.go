package daos

import (
	"github.com/constant-money/constant-web/event/models"
	"github.com/jinzhu/gorm"
)

type TxDAO struct{}

func (t *TxDAO) GetAllPending() ([]models.Tx, error) {
	txs := []models.Tx{}
	err := models.Database().Where("hash != -1 and status = -1").Find(&txs).Error
	if err != nil {
		return nil, err
	}
	return txs, nil
}

func (t *TxDAO) GetTxPending(query string) ([]models.Tx, error) {
	txs := []models.Tx{}
	err := models.Database().Where(query).Find(&txs).Error
	if err != nil {
		return nil, err
	}
	return txs, nil
}

func (t *TxDAO) GetByHash(hash string) (models.Tx, error) {
	tx := models.Tx{}
	err := models.Database().Where("hash = ?", hash).First(&tx).Error
	return tx, err
}

func (t *TxDAO) New(tx *models.Tx) error {
	err := models.Database().Create(tx).Error
	return err
}

func (t *TxDAO) Update(tx *models.Tx) error {
	err := models.Database().Save(tx).Error
	return err
}

func (t *TxDAO) NewWithTnx(tx *models.Tx, tnx *gorm.DB) error {
	err := tnx.Create(tx).Error
	return err
}
