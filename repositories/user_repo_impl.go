package repositories

import (
	"github.com/iqbalmahad/todolist-app-golang.git/models"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	db *gorm.DB // GORM database connection
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserRepoImpl{db: db}
}

func (ur *UserRepoImpl) Create(user models.User) error {
	return ur.db.Create(&user).Error
}

func (ur *UserRepoImpl) Read(id string) (models.User, error) {
	var user models.User
	if err := ur.db.Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (ur *UserRepoImpl) Update(user models.User) error {
	return ur.db.Model(&user).Updates(user).Error
}

func (ur *UserRepoImpl) Delete(id string) error {
	return ur.db.Where("id = ?", id).Delete(&models.User{}).Error
}

func (ur *UserRepoImpl) Index() ([]models.User, error) {
	var users []models.User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

