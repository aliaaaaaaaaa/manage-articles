package api

import (
	"github.com/labstack/echo"
	"manageArticles_/internal/model"
	"manageArticles_/pkg/utils"
	"net/http"
)

func (h *handler) login(c echo.Context) error {
	user := new(model.User)
	err := c.Bind(user)
	if err != nil {
		return echo.ErrUnauthorized
	}
	user, err = h.userRepo.FindUser(user.Email, user.Password)
	if err != nil {
		return err
	}

	token, err := utils.GetToken(*user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})

}
