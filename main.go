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
	todoRepository := repository.NewTodoRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	todoUsecase := usecase.NewTodoUsecase(todoRepository)
	userController := controller.NewUserController(userUsecase)
	todoController := controller.NewTodoController(todoUsecase)
	e := router.NewRouter(userController, todoController)
	e.Logger.Fatal(e.Start(":8080"))
}