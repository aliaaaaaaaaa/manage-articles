package api

import (
	"github.com/labstack/echo"
	"manageArticles_/pkg/utils"
	"net/http"
)

func (h *handler) getArticle(c echo.Context) error {
	limit := utils.StringToInt(c.Param("limit"))
	offset := utils.StringToInt(c.Param("offset"))

	articles := h.ArticleRepo.ShowArticles(limit, offset)
	return c.JSON(http.StatusOK, articles)
}
