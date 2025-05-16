package router

import (
	"github.com/shinsx/golang-blog/controller"
	"net/http"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.POST("/login", func(c echo.Context) error {
			user, err := uc.LogIn(c)
			if err != nil {
					return err // uc.LogIn 側で c.JSON(...) していればそれが返る
			}
			return c.JSON(http.StatusOK, user)
	})
	return e
}
