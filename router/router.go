package router

import (
	"go-todo-api/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, tc controller.ITodoController) *echo.Echo{
	e := echo.New()
	//エンドポイントの追加
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	t := e.Group("/todos")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllTodos)
	t.GET("/:todoId", tc.GetTodoById)
	t.POST("", tc.CreateTodo)
	t.PUT("/:todoId", tc.UpdateTodo)
	t.DELETE("/:todoId", tc.DeleteTodo)
	return e
}