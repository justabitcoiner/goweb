package handlers

import (
	"goweb/db"
	"goweb/views"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func CreateSession(c echo.Context, userId int) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["userId"] = userId
	err = sess.Save(c.Request(), c.Response())
	return err
}

// Sign Up
func GetSignUpView(c echo.Context) error {
	return views.SignUp().Render(c.Request().Context(), c.Response())
}

func SignUp(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	err := db.SignUp(email, password)
	if err != nil {
		return c.String(422, err.Error())
	}

	c.Response().Header().Set("HX-Redirect", "/signin")
	return c.NoContent(200)
}

// Sign In
func GetSignInView(c echo.Context) error {
	return views.SignIn().Render(c.Request().Context(), c.Response())
}

func SignIn(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	userId, err := db.SignIn(email, password)
	if err != nil {
		return c.String(422, err.Error())
	}

	err = CreateSession(c, userId)
	if err != nil {
		return c.String(422, err.Error())
	}

	c.Response().Header().Set("HX-Redirect", "/")
	return c.NoContent(200)
}
