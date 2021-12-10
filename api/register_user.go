package api

import (
	"github.com/labstack/echo"
	"manageArticles_/internal/model"
	"manageArticles_/pkg/utils"
	"net/http"
)

func (h *handler) registerUser(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return err
	}
	createUser, err := h.userRepo.CreateUser(&user)
	token, err := utils.GetToken(*createUser)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})

}
