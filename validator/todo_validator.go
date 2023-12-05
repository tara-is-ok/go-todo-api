package validator

import (
	"go-todo-api/models"

	validation "github.com/go-ozzo/ozzo-validation"
)




type ITodoValidator interface {
	TodoValidate(todo models.Todo) error
}

type todoValidator struct {}

func NewTodoValidator() ITodoValidator {
	return &todoValidator{}
}


func (tv *todoValidator) TodoValidate(todo models.Todo) error {
	return validation.ValidateStruct(&todo, validation.Field(
		&todo.Title,
		validation.Required.Error("title is required"),
		validation.RuneLength(1,10).Error("limit 10 charactor"),
	))
}