package main

import (
	"log"
	"manageArticles_/api"
	"manageArticles_/config"
	"manageArticles_/internal/repo"
	"manageArticles_/pkg/db"
)

func main() {
	config := config.LoadConfig()
	gorm := db.GetGorm(config)
	userRepo := repo.NewUserRepositoryImpl(gorm)
	articleRepo := repo.NewArticleRepositoryImpl(gorm)
	handler := api.NewHandler(userRepo, articleRepo, config)
	log.Println("starting the web server")
	handler.StartWebServer()
	select {}
}
