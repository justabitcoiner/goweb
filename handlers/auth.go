package handlers

import (
	"goweb/db"
	"goweb/views"

	"github.com/labstack/echo/v4"
)

// Sign Up
func GetSignUpView(c echo.Context) error {
	return views.SignUp().Render(c.Request().Context(), c.Response())
}

func SignUp(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	err := db.SignUp(email, password)
	if err != nil {
		return err
	}
	return nil
}

// Sign In
func GetSignInView(c echo.Context) error {
	return views.SignIn().Render(c.Request().Context(), c.Response())
}

func SignIn(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	err := db.SignIn(email, password)
	if err != nil {
		return err
	}
	return nil
}
