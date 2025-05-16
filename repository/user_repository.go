package repository

import (
	"github.com/shinsx/golang-blog/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) (*model.User, error) {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
