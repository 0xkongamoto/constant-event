package daos

import (
	"github.com/constant-money/constant-event/models"
	wm "github.com/constant-money/constant-web-api/models"
	"github.com/jinzhu/gorm"
)

type TxDAO struct{}

func (t *TxDAO) GetAllPending() ([]wm.Tx, error) {
	txs := []wm.Tx{}
	err := models.Database().Where("hash != -1 and status = -1").Find(&txs).Error
	if err != nil {
		return nil, err
	}
	return txs, nil
}

func (t *TxDAO) GetTxPending(query string) ([]wm.Tx, error) {
	txs := []wm.Tx{}
	err := models.Database().Where(query).Find(&txs).Error
	if err != nil {
		return nil, err
	}
	return txs, nil
}

func (t *TxDAO) GetByHash(hash string) (wm.Tx, error) {
	tx := wm.Tx{}
	err := models.Database().Where("hash = ?", hash).First(&tx).Error
	return tx, err
}

func (t *TxDAO) New(tx *wm.Tx) error {
	err := models.Database().Create(tx).Error
	return err
}

func (t *TxDAO) Update(tx *wm.Tx) error {
	err := models.Database().Save(tx).Error
	return err
}

func (t *TxDAO) NewWithTnx(tx *wm.Tx, tnx *gorm.DB) error {
	err := tnx.Create(tx).Error
	return err
}
