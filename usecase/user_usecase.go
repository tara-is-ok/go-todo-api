package usecase

import (
	"go-todo-api/models"
	"go-todo-api/repository"
	"go-todo-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	//値型(models.User)を取ることで、このメソッド内で新しいユーザー情報を変更し、データベースに挿入
	SignUp(user models.User) (models.UserResponse,error)
	Login(user models.User)(string, error) //JWTトークンを返す
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase{
	return &userUsecase{ur, uv}
}

func (uu *userUsecase)SignUp(user models.User)(models.UserResponse,error){
	if err := uu.uv.UserValidate(user); err != nil {
		return models.UserResponse{}, err
	}
	//パスワードをハッシュ化
	hash, err := 	bcrypt.GenerateFromPassword([]byte(user.Password),10) //第2引数は暗号の複雑さ
	if err != nil {
		return models.UserResponse{},err
	}
	newUser := models.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil{
		return models.UserResponse{}, err
	}
	resUser := models.UserResponse{
		ID: newUser.ID,
		Email: newUser.Email,
	}
	return resUser,nil
}

func (uu *userUsecase) Login(user models.User)(string, error){
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}
	//clientからくるemailがdbに存在するか確認する
	storedUser := models.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil{
		return "", err
	}
	//パスワードの一致確認
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	//JWTトークンの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp": time.Now().Add(time.Hour * 12).Unix(), //有効期限
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil{
		return "",err
	}
	return tokenString, nil
}