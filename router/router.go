package router

import (
	"go-todo-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.ITodoController) *echo.Echo{
	e := echo.New()
	//corsのmiddleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//アクセスを許可するフロントエンドのドメイン
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		//許可するヘッダー一覧
		//ヘッダー経由でCSRFトークンを受け取る
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,echo.HeaderAccessControlAllowHeaders,echo.HeaderXCSRFToken},
		AllowMethods: []string{"GET","PUT","POST","DELETE"},
		//cookieの送受信を可能にする
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath: "/",
		CookieDomain: os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		//postmanなどで動作確認をする際にfalseにする必要になる
		CookieSameSite:http.SameSiteNoneMode ,
		// CookieSameSite: http.SameSiteDefaultMode,
		//有効期限
		//CookieMaxAge: 60
	}))
	//エンドポイントの追加
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	e.GET("/csrf", uc.CsrfToken)
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