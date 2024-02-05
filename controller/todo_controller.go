package controller

import (
	"go-todo-api/models"
	"go-todo-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ITodoController interface {
	GetAllTodos(c echo.Context) error
	GetTodoById(c echo.Context) error
	CreateTodo(c echo.Context) error
	UpdateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

type todoController struct {
	tu usecase.ITodoUsecase
}

func NewTodoController(tu usecase.ITodoUsecase) ITodoController {
	return &todoController{tu}
}

func (tc *todoController) GetAllTodos(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	tags := c.QueryParams()["tag"]
	if len(tags) != 0 {
		todosRes, err := tc.tu.GetTodosByTags(tags)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, todosRes)
	}
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

	// リクエストボディを一度マップにバインド
	var requestBody map[string]interface{}
	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tagsBody, ok := requestBody["tags"].([]interface{})
	if !ok {
		return c.JSON(http.StatusBadRequest, "tags must be array type")
	}
	var tags []models.Tag
	for _, tag := range tagsBody {
		tagName, ok := tag.(string)
		if !ok {
			return c.JSON(http.StatusBadRequest, "tag must be string type")
		}
		tagStruct := models.Tag{
			Name: tagName,
		}
		tags = append(tags, tagStruct)
	}

	todo := models.Todo{}
	todo.Title, ok = requestBody["title"].(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, "title is required")
	}
	todo.Tags = tags
	todo.UserId = uint(userId.(float64))
	todoRes, err := tc.tu.CreateTodo(todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, todoRes)
}

func (tc *todoController) UpdateTodo(c echo.Context) error {
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
