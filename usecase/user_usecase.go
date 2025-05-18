package usecase

import (
	"github.com/shinsx/golang-blog/model"
	"github.com/shinsx/golang-blog/repository"

	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type IUserUsecase interface {
	Login(user model.User) (*model.User, string, error)
	SignUp(user model.User) (model.UserResponse, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	newUser := model.User{Name: user.Name, Email: user.Email, Password: string(user.Password)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (*model.User, string, error) {
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return nil, "", err
	}
	if storedUser.Password != user.Password {
		return nil, "", errors.New("password is incorrect")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return nil, "", err
	}
	return &storedUser, tokenString, nil
}
