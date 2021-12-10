package api

import (
	"github.com/dgrijalva/jwt-go"
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
	user := c.Get("user").(*jwt.Token)
	Claims := user.Claims.(utils.JwtCustomClaims)
	createArticle, err := h.ArticleRepo.CreateArticle(article, model.User{Id: Claims.UserId})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, createArticle)
}
