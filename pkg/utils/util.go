package utils

import (
	"github.com/dgrijalva/jwt-go"
	"manageArticles_/internal/model"
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
