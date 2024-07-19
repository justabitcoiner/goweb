package main

import (
	"goweb/db"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	db.Connect()
	defer db.Disconnect()

	app.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello world")
	})

	app.Logger.Fatal(app.Start("localhost:9000"))
}
