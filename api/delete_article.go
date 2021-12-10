package api

import (
	"github.com/labstack/echo"
	"manageArticles_/pkg/utils"
	"net/http"
)

func (h *handler) deleteArticle(c echo.Context) error {
	id := utils.StringToInt(c.Param("id"))
	userId := utils.GetUserID(c)
	err := h.articleService.DeleteArticle(id, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}
