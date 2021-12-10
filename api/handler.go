package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"manageArticles_/config"
	"manageArticles_/internal/repo"
	"manageArticles_/pkg/utils"
)

type handler struct {
	userRepo    repo.UserRepository
	config      *config.ManageArticalConfig
	ArticleRepo repo.ArticleRepository
}

func NewHandler(userRepo repo.UserRepository, articleRepo repo.ArticleRepository, config *config.ManageArticalConfig) *handler {
	return &handler{userRepo: userRepo, config: config, ArticleRepo: articleRepo}
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
		SigningKey: []byte("lololo"),
	}
	Requests.Use(middleware.JWTWithConfig(configEcho))

	Requests.GET("/article/:limit/:offset", h.getArticle)
	Requests.POST("/article", h.createArticle)
	Requests.DELETE("/article/:id", h.deleteArticle)
	Requests.PUT("/article", h.updateArticle)
	echoServer.Logger.Fatal(echoServer.Start(":" + h.config.ServerConfig.Port))
}
