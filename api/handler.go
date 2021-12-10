package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"manageArticles_/config"
	"manageArticles_/internal/repo"
	"manageArticles_/pkg/utils"
)

type handler struct {
	userRepo repo.UserRepository
	config   config.ManageArticalConfig
}

func newHandler(userRepo repo.UserRepository) *handler {
	return &handler{userRepo: userRepo}
}

func (h handler) StartWebServer() {
	echoServer := echo.New()
	echoServer.POST("/login", h.login)
	echoServer.POST("/register", h.registerUser)

	echoServer.Use(middleware.Recover())
	echoServer.Use(middleware.CORS())
	Requests := echoServer.Group("")
	configEcho := middleware.JWTConfig{
		Claims:     &utils.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	Requests.Use(middleware.JWTWithConfig(configEcho))
	echoServer.Logger.Fatal(echoServer.Start(":" + h.config.ServerConfig.Port))
}
