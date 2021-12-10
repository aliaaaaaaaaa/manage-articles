package api

import (
	"github.com/labstack/echo"
	"manageArticles_/internal/model"
	"manageArticles_/pkg/utils"
	"net/http"
)

func (h *handler) updateArticle(c echo.Context) error {
	var article model.Article
	id := utils.StringToInt(c.Param("id"))
	err := c.Bind(&article)
	article.Id = id
	if err != nil {
		return err
	}
	userId := utils.GetUserID(c)
	err = h.articleService.UpdateArticle(&article, userId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, article)
}
