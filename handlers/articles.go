package handlers

import (
	"fmt"
	"goweb/db"
	"goweb/views"

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
