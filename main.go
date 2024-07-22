package main

import (
	"fmt"
	"goweb/db"
	"goweb/handlers"
	"goweb/middlewares"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	db.Connect()
	defer db.Disconnect()

	// Static
	app.Static("static", "static")

	// Middlewares
	app.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	app.GET("/", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		userId := sess.Values["userId"].(int)
		return c.String(200, fmt.Sprintf("Hello world, user: %v", userId))
	}, middlewares.AuthMiddleware)
	app.GET("/signup", handlers.GetSignUpView)
	app.POST("/signup", handlers.SignUp)
	app.GET("/signin", handlers.GetSignInView)
	app.POST("/signin", handlers.SignIn)
	app.GET("/articles", handlers.GetArticleListView)
	app.GET("/articles/:id", handlers.GetArticleDetailView)
	app.GET("/articles/new", handlers.GetArticleNew, middlewares.AuthMiddleware)
	app.POST("/articles/new", handlers.GetArticleNew, middlewares.AuthMiddleware)
	app.Logger.Fatal(app.Start("localhost:9000"))
}
