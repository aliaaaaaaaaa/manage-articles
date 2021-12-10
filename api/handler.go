package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"manageArticles_/config"
	"manageArticles_/internal/service"
)

type handler struct {
	config         *config.ManageArticalConfig
	userService    *service.UserService
	articleService *service.ArticleService
}

func NewHandler(userService *service.UserService, articleService *service.ArticleService, config *config.ManageArticalConfig) *handler {
	return &handler{
		config:         config,
		userService:    userService,
		articleService: articleService,
	}
}

func (h handler) StartWebServer() {
	echoServer := echo.New()
	echoServer.POST("/login", h.login)
	echoServer.POST("/register", h.registerUser)

	echoServer.Use(middleware.Recover())
	echoServer.Use(middleware.CORS())
	Requests := echoServer.Group("")
	configEcho := middleware.JWTConfig{
		Claims:     &service.JwtCustomClaims{},
		SigningKey: []byte("lololo"),
	}
	Requests.Use(middleware.JWTWithConfig(configEcho))

	Requests.GET("/article/:limit/:offset", h.getArticle)
	Requests.POST("/article", h.createArticle)
	Requests.DELETE("/article/:id", h.deleteArticle)
	Requests.PUT("/article/:id", h.updateArticle)
	echoServer.Logger.Fatal(echoServer.Start(":" + h.config.ServerConfig.Port))
}
