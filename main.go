package main

import (
	"goweb/db"
	"goweb/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	db.Connect()
	defer db.Disconnect()

	// Static
	app.Static("static", "static")

	app.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello world")
	})
	app.GET("/signup", handlers.GetSignUpView)
	app.POST("/signup", handlers.SignUp)
	app.GET("/signin", handlers.GetSignInView)
	app.POST("/signin", handlers.SignIn)
	app.Logger.Fatal(app.Start("localhost:9000"))
}
