package daos

import (
	"time"

	"github.com/constant-money/constant-event/models"
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
func (ud *UserDAO) GetAllUserWalletPending() ([]models.UserWallet, error) {
	userWallets := []models.UserWallet{}
	err := models.Database().Where("status != ? AND expired_at > ?", models.UserWalletStatusDone, time.Now().UnixNano()/int64(time.Second)).Find(&userWallets).Error
	return userWallets, err
}

// GetAllUsersNeedCheckKYC : ...
func (ud *UserDAO) GetAllUsersNeedCheckKYC() (*[]models.User, error) {
	users := []models.User{}
	err := models.Database().Where("(verified_level = 4 OR verified_level = 5) AND primetrust_contact_id <> '' ").Find(&users).Error
	return &users, err
}

// Update : user wallet
func (ud *UserDAO) Update(uw *models.UserWallet) error {
	err := models.Database().Save(uw).Error
	return err
}

// UpdateVerifiedLevel : user
func (ud *UserDAO) UpdateVerifiedLevel(u *models.User) error {
	err := models.Database().Save(u).Error
	return err
}
