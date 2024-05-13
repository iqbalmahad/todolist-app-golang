package repositories

import (
	"github.com/iqbalmahad/todolist-app-golang.git/models"
	"gorm.io/gorm"
)

type authRepoImpl struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &authRepoImpl{
		db: db,
	}
}

func (ar *authRepoImpl) GetByUsernameOrEmail(usernameOrEmail string) (*models.User, error) {
	var user models.User
	if err := ar.db.Where("username = ? OR email = ?", usernameOrEmail, usernameOrEmail).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
