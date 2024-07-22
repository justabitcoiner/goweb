package handlers

import (
	"fmt"
	"goweb/db"
	"goweb/views"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetArticleListView(c echo.Context) error {
	article_list, err := db.GetArticleList()
	if err != nil {
		return fmt.Errorf("cannot get article list")
	}
	return views.ArticleList(article_list).Render(c.Request().Context(), c.Response())
}

func GetArticleDetailView(c echo.Context) error {
	id := c.Param("id")
	article, err := db.GetArticleDetail(id)
	if err != nil {
		return fmt.Errorf("cannot get article detail")
	}

	return views.ArticleDetail(*article).Render(c.Request().Context(), c.Response())
}

func GetArticleNew(c echo.Context) error {
	if c.Request().Method == "GET" {
		return views.ArticleNew().Render(c.Request().Context(), c.Response())
	} else {
		sess, _ := session.Get("session", c)
		userId := sess.Values["userId"].(int)
		title := c.FormValue("title")
		content := c.FormValue("content")

		err := db.CreateNewArticle(userId, title, content)
		if err != nil {
			return c.String(422, "cannot create new article")
		}
	}

	c.Response().Header().Set("HX-Redirect", "/articles")
	return c.String(200, "create new article success")
}
