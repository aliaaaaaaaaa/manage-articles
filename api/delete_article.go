package api

import (
	"fmt"
	"github.com/labstack/echo"
	"manageArticles_/pkg/utils"
	"net/http"
)

func (h *handler) deleteArticle(c echo.Context) error {
	id := utils.StringToInt(c.Param("id"))
	fmt.Println(id)
	err := h.ArticleRepo.DeleteArticle(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
}
