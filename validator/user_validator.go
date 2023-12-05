package validator

import (
	"go-todo-api/models"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)




type IUserValidator interface {
	UserValidate(user models.User) error
}

type userValidator struct {}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}


func (uv *userValidator) UserValidate(user models.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1,30).Error("limit 30 charactor"),
			is.Email.Error("is not valid format"),
		),validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(10,30).Error("min 10 max 30 charactor"),
		),
	)
}