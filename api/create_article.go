package api

import (
	"github.com/labstack/echo"
	"manageArticles_/internal/model"
	"manageArticles_/pkg/utils"
	"net/http"
)

func (h *handler) createArticle(c echo.Context) error {
	var article model.Article
	err := c.Bind(&article)
	if err != nil {
		return err
	}
	userId := utils.GetUserID(c)
	err = h.articleService.CreateArticle(&article, userId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, article)
}
