package api

import (
	"github.com/labstack/echo"
	"manageArticles_/internal/model"
	"net/http"
)

func (h *handler) updateArticle(c echo.Context) error {
	var article model.Article
	err := c.Bind(&article)
	if err != nil {
		return err
	}
	updateArticle, err := h.ArticleRepo.UpdateArticle(article)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, updateArticle)
}
