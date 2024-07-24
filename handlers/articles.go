package handlers

import (
	"goweb/db"
	"goweb/views"
	"strconv"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GetArticleListView(c echo.Context) error {
	article_list, err := db.GetArticleList()
	if err != nil {
		return c.String(422, "cannot get article list")
	}
	return views.ArticleList(article_list).Render(c.Request().Context(), c.Response())
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

func GetArticleEditView(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(422, "id is invalid")
	}

	switch c.Request().Method {
	case "GET":
		mode := c.QueryParam("mode")
		if mode == "" {
			article, err := db.GetArticleDetail(id)
			if err != nil {
				return c.String(422, "cannot get article detail")
			}

			return views.ArticleDetail(*article).Render(c.Request().Context(), c.Response())
		} else if mode == "edit" {
			userId, err := GetCurrentUserId(c)
			if err != nil {
				return c.String(422, err.Error())
			}
			article, err := db.GetArticleDetail(id)
			if err != nil {
				return c.String(422, "cannot get article edit view")
			}
			if article.UserId != userId {
				return views.Forbidden_403().Render(c.Request().Context(), c.Response())
			}

			return views.ArticleEdit(*article).Render(c.Request().Context(), c.Response())
		} else {
			return c.String(422, "does not support this view mode")
		}
	case "PATCH":
		userId, err := GetCurrentUserId(c)
		if err != nil {
			return c.String(422, err.Error())
		}
		title := c.FormValue("title")
		content := c.FormValue("content")

		err = db.UpdateArticle(userId, id, title, content)
		if err != nil {
			return c.String(422, "cannot update article")
		}

		c.Response().Header().Set("HX-Redirect", "/articles")
		return c.String(200, "update article success")
	case "DELETE":
		userId, err := GetCurrentUserId(c)
		if err != nil {
			return c.String(422, err.Error())
		}
		err = db.DeleteArticle(userId, id)
		if err != nil {
			return c.String(422, "cannot delete article")
		}

		c.Response().Header().Set("HX-Redirect", "/articles")
		return c.String(200, "delete article success")
	default:
		return c.String(405, "Method not support")
	}
}
