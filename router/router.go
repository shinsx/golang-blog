package router

import (
	"github.com/shinsx/golang-blog/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.POST("/login", uc.LogIn)
	return e
}
