package router

import (
	"github.com/shinsx/golang-blog/controller"

	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.GET("/", Hello)
	e.POST("/login", uc.LogIn)
	e.POST("/signup", uc.SignUp)
	t := e.Group("/articles")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", LogInCheck)
	return e
}

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func LogInCheck(c echo.Context) error {
	return c.String(http.StatusOK, "ログインしてる")
}
