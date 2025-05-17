package usecase

import (
	"github.com/shinsx/golang-blog/model"
	"github.com/shinsx/golang-blog/repository"

	"errors"
)

type IUserUsecase interface {
	Login(user model.User) (*model.User, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) Login(user model.User) (*model.User, error) {
	storedUser := model.User{}
	getUser, err := uu.ur.GetUserByEmail(&storedUser, user.Email)
	if err != nil {
		return nil, err
	}
	if storedUser.Password != user.Password {
		return nil, errors.New("password is incorrect")
	}
	return getUser, nil
}
