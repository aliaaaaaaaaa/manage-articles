package api

import (
	"github.com/labstack/echo"
	"manageArticles_/internal/model"
	"manageArticles_/pkg/utils"
	"net/http"
)

func (h *handler) updateArticle(c echo.Context) error {
	var article model.Article
	err := c.Bind(&article)
	if err != nil {
		return err
	}
	id := utils.GetUser(c)
	FindedUser, err := h.userRepo.FindBtID(id)
	article.Author = FindedUser.Name
	article.UserID = FindedUser.Id
	updateArticle, err := h.ArticleRepo.UpdateArticle(article)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, updateArticle)
}
