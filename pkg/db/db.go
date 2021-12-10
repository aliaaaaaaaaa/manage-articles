package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"manageArticles_/config"
)

func GetGorm(config *config.ManageArticalConfig) *gorm.DB {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=localhost port=5432 user=%s dbname=%s password=%s sslmode=disable", config.DbConfig.DbUser, config.DbConfig.DbName, config.DbConfig.DbPass))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
