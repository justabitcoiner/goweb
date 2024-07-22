package middlewares

import (
	"goweb/views"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil || sess.Values["userId"] == nil {
			return views.Unauthorized().Render(c.Request().Context(), c.Response())
		}
		next(c)
		return nil
	}
}
