package controller

import (
	"github.com/shinsx/golang-blog/model"
	"github.com/shinsx/golang-blog/usecase"

	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	LogIn(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) LogIn(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	getUser, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, getUser)
}
