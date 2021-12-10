package api

import (
	"github.com/labstack/echo"
	"manageArticles_/pkg/utils"
	"net/http"
)

func (h *handler) getArticle(c echo.Context) error {
	limit := utils.StringToInt(c.Param("limit"))
	offset := utils.StringToInt(c.Param("offset"))
	userId := utils.GetUserID(c)
	articles := h.articleService.GetAllArticle(limit, offset, userId)
	return c.JSON(http.StatusOK, articles)
}
