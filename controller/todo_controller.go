package controller

import (
	"go-todo-api/models"
	"go-todo-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type  ITodoController interface {
	GetAllTodos(c echo.Context) error
	GetTodoById(c echo.Context) error
	CreateTodo(c echo.Context)error
	UpdateTodo(c echo.Context)error
	DeleteTodo(c echo.Context)error
}

type todoController struct {
	tu usecase.ITodoUsecase
}

func NewTodoController(tu usecase.ITodoUsecase) ITodoController{
	return &todoController{tu}
}

func (tc *todoController) GetAllTodos(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	//userIdがanyのため型変換を行う
	todosRes, err := tc.tu.GetAllTodos(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todosRes)
}

func (tc *todoController) GetTodoById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("todoId")
	todoId, _ := strconv.Atoi(id)
	todoRes, err := tc.tu.GetTodoById(uint(userId.(float64)), uint(todoId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todoRes)
}

func (tc *todoController) CreateTodo(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	todo := models.Todo{}
	//request bodyをtodo structに代入
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	todo.UserId = uint(userId.(float64))
	todoRes, err := tc.tu.CreateTodo(todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, todoRes)
}

func (tc *todoController) UpdateTodo(c echo.Context)error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("todoId")
	todoId, _ := strconv.Atoi(id)

	todo := models.Todo{}	
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	todoRes, err := tc.tu.UpdateTodo(todo, uint(userId.(float64)), uint(todoId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, todoRes)
}

func (tc *todoController) DeleteTodo(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("todoId")
	todoId, _ := strconv.Atoi(id)

	err := tc.tu.DeleteTodo(uint(userId.(float64)), uint(todoId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}