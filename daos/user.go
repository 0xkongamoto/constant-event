package daos

import (
	"strings"

	"github.com/constant-money/constant-event/models"
	wm "github.com/constant-money/constant-web-api/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
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

// AllUserWallets : metadata
func (ud *UserDAO) AllUserWallets(metadata string) ([]*wm.UserWallet, error) {
	var userWallets []*wm.UserWallet
	if err := models.Database().Where("lower(metadata) = ? and source = ethereum", strings.ToLower(metadata)).Find(&userWallets).Error; err != nil {
		return nil, errors.Wrap(err, "db.Where.Find")
	}
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
