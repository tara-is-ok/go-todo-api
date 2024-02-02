package usecase

import (
	"go-todo-api/models"
	"go-todo-api/repository"
	"go-todo-api/validator"
)

type ITodoUsecase interface {
	GetAllTodos(userId uint) ([]models.TodoResponse, error)
	GetTodoById(userId uint, todoId uint) (models.TodoResponse, error)
	CreateTodo(todo models.Todo) (models.TodoResponse, error)
	UpdateTodo(todo models.Todo, userId uint, todoId uint) (models.TodoResponse, error)
	DeleteTodo(userId uint, todoId uint) error
}

type todoUsecase struct {
	tr repository.ITodoRepository
	tv validator.ITodoValidator
}

func NewTodoUsecase(tr repository.ITodoRepository, tv validator.ITodoValidator) ITodoUsecase {
	return &todoUsecase{tr, tv}
}

func (tu *todoUsecase) GetAllTodos(userId uint) ([]models.TodoResponse, error) {
	todos := []models.Todo{}
	if err := tu.tr.GetAllTodos(&todos, userId); err != nil {
		return nil, err
	}
	resTodos := []models.TodoResponse{}
	for _, v := range todos {
		t := models.TodoResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Tags:      v.Tags,
		}
		resTodos = append(resTodos, t)
	}
	return resTodos, nil
}

func (tu *todoUsecase) GetTodoById(userId uint, todoId uint) (models.TodoResponse, error) {
	todo := models.Todo{}
	if err := tu.tr.GetTodoById(&todo, userId, todoId); err != nil {
		return models.TodoResponse{}, err
	}
	resTodo := models.TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
		Tags:      todo.Tags,
	}
	return resTodo, nil
}

func (tu *todoUsecase) CreateTodo(todo models.Todo) (models.TodoResponse, error) {
	if err := tu.tv.TodoValidate(todo); err != nil {
		return models.TodoResponse{}, err
	}

	if err := tu.tr.CreateTodo(&todo); err != nil {
		return models.TodoResponse{}, err
	}
	resTodo := models.TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
		Tags:      todo.Tags,
	}
	return resTodo, nil
}

func (tu *todoUsecase) UpdateTodo(todo models.Todo, userId uint, todoId uint) (models.TodoResponse, error) {
	if err := tu.tr.UpdateTodo(&todo, userId, todoId); err != nil {
		return models.TodoResponse{}, err
	}
	resTodo := models.TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
		Tags:      todo.Tags,
	}
	return resTodo, nil
}

func (tu *todoUsecase) DeleteTodo(userId uint, todoId uint) error {
	if err := tu.tr.DeleteTodo(userId, todoId); err != nil {
		return err
	}
	return nil
}
