package daos

import (
	"github.com/constant-money/constant-event/models"
	wm "github.com/constant-money/constant-web-api/models"
	"github.com/jinzhu/gorm"
)

// UserDAO :
type UserDAO struct {
	db *gorm.DB
}

// InitUserDAO :
func InitUserDAO(database *gorm.DB) *UserDAO {
	return &UserDAO{
		db: database,
	}
}

// GetAllUserWalletPending : ...
func (ud *UserDAO) GetAllUserWalletPending() ([]wm.UserWallet, error) {
	userWallets := []wm.UserWallet{}
	return userWallets, nil
}

// GetAllUsersNeedCheckKYC : ...
func (ud *UserDAO) GetAllUsersNeedCheckKYC() (*[]wm.User, error) {
	users := []wm.User{}
	err := models.Database().Where("(verified_level = 4 OR verified_level = 5) AND primetrust_contact_id <> '' ").Find(&users).Error
	return &users, err
}

// Update : user wallet
func (ud *UserDAO) Update(uw *wm.UserWallet) error {
	err := models.Database().Save(uw).Error
	return err
}

// UpdateVerifiedLevel : user
func (ud *UserDAO) UpdateVerifiedLevel(u *wm.User) error {
	err := models.Database().Save(u).Error
	return err
}
