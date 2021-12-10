package service

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"manageArticles_/internal/model"
	"manageArticles_/internal/repo"
	"time"
)

type UserService struct {
	userRepo repo.UserRepository
}
type JwtCustomClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

func NewUserService(userRepo repo.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) GetToken(user model.User) (string, error) {
	claims := &JwtCustomClaims{
		user.Id,
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
func (u *UserService) Login(username string, password string) (string, error) {
	password = u.returnUserEncryptedPassword(username, password)
	user, err := u.userRepo.FindUser(username, password)
	if err != nil {
		return "", echo.ErrUnauthorized
	}
	token, err := u.GetToken(*user)
	return token, err
}
func (u *UserService) Register(user model.User) (string, error) {
	user.Password = u.returnUserEncryptedPassword(user.Email, user.Password)
	user.IsActive = true
	err := u.userRepo.CreateUser(&user)
	if err != nil {
		return "", err
	}
	token, err := u.GetToken(user)
	return token, err
}

func (u *UserService) returnUserEncryptedPassword(username string, password string) string {
	shaInfo := sha256.Sum256([]byte(fmt.Sprintf("%xlol%x", sha256.Sum256([]byte(password)), sha256.Sum256([]byte(username)))))
	return base64.URLEncoding.EncodeToString([]byte(shaInfo[:]))
}
