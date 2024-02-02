package repository

import (
	"fmt"
	"go-todo-api/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITodoRepository interface {
	GetAllTodos(todos *[]models.Todo, userId uint) error
	GetTodoById(todo *models.Todo, userId uint, todoId uint) error
	CreateTodo(todo *models.Todo) error
	UpdateTodo(todo *models.Todo, userId uint, todoId uint) error
	DeleteTodo(userId uint, todoId uint) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) ITodoRepository {
	return &todoRepository{db}
}

func (tr *todoRepository) GetAllTodos(todos *[]models.Todo, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(todos).Preload("Tags").Find(todos).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepository) GetTodoById(todo *models.Todo, userId uint, todoId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(todo, todoId).Preload("Tags").Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepository) CreateTodo(todo *models.Todo) error {
	if err := tr.db.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepository) UpdateTodo(todo *models.Todo, userId uint, todoId uint) error {
	fmt.Println("渡ってきた時", todo.Tags)
	result := tr.db.Model(todo).Clauses(clause.Returning{}).Where("id=? AND user_id=?", todoId, userId).Updates(models.Todo{Title: todo.Title, Tags: todo.Tags})
	fmt.Println("Updateあと", todo.Tags)
	tr.db.Model(todo).Association("Tags").Replace(todo.Tags)
	fmt.Println("アソシエーション後", todo.Tags)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *todoRepository) DeleteTodo(userId uint, todoId uint) error {
	result := tr.db.Where("id=? AND user_id=?", todoId, userId).Delete(&models.Todo{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
