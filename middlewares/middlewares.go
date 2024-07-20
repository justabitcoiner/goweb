package middlewares

import (
	"goweb/views"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := session.Get("session", c)
		if err != nil {
			return views.Unauthorized().Render(c.Request().Context(), c.Response())
		}
		next(c)
		return nil
	}
}
