package api

import (
	"github.com/labstack/echo"
	"manageArticles_/internal/model"
	"net/http"
)

func (h *handler) registerUser(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return err
	}
	token, err := h.userService.Register(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})

}
