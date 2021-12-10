package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"manageArticles_/internal/service"
	"strconv"
)

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func GetUserID(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	Claims := user.Claims.(*service.JwtCustomClaims)
	return Claims.UserId
}
