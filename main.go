package main

import (
	"go-todo-api/controller"
	"go-todo-api/db"
	"go-todo-api/repository"
	"go-todo-api/router"
	"go-todo-api/usecase"
)

func main(){
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}