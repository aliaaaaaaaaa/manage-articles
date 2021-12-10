package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"manageArticles_/internal/model"
	"strconv"
	"time"
)

type JwtCustomClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

func GetToken(u model.User) (string, error) {
	claims := &JwtCustomClaims{
		u.Id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("lololo"))
	if err != nil {
		return "", err
	}
	return t, nil
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
func ReturnUserEncryptedPassword(username string, password string) string {
	shaInfo := sha256.Sum256([]byte(fmt.Sprintf("%xlol%x", sha256.Sum256([]byte(password)), sha256.Sum256([]byte(username)))))
	return base64.URLEncoding.EncodeToString([]byte(shaInfo[:]))
}

func GetUser(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	Claims := user.Claims.(*JwtCustomClaims)
	return Claims.UserId
}
