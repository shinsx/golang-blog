package main

import (
	"github.com/shinsx/golang-blog/controller"
	"github.com/shinsx/golang-blog/db"
	"github.com/shinsx/golang-blog/repository"
	"github.com/shinsx/golang-blog/router"
	"github.com/shinsx/golang-blog/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":3000"))
}
