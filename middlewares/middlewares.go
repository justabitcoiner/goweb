package middlewares

import (
	"context"
	"goweb/handlers"
	"goweb/views"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		mode := c.QueryParam("mode")
		if c.Request().Method == "GET" && mode != "edit" {
			return next(c)
		}

		sess, err := session.Get("session", c)
		if err == nil {
			if sess.Values["userId"] != nil {
				return next(c)
			}
		}

		return views.Unauthorized_401().Render(c.Request().Context(), c.Response())
	}
}

func WithContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, _ := handlers.GetCurrentUserId(c)
		ctx := context.WithValue(context.Background(), "userId", userId)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
