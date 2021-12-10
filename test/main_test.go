package test

import (
	"github.com/jinzhu/gorm"
	"manageArticles_/internal/model"
	"manageArticles_/internal/repo"
	"os"
	"testing"
)

var userRepo repo.UserRepository
var articleRepo repo.ArticleRepository

func TestMain(m *testing.M) {
	setup()

	r := m.Run()
	os.Exit(r)
}
func setup() {
	dsn := "host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/tehran"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return
	}
	db.Exec("drop table if exists users")
	db.Exec("drop table if exists articles")
	db.AutoMigrate(&model.User{}, &model.Article{})
	userRepo = repo.NewUserRepositoryImpl(db)
	articleRepo = repo.NewArticleRepositoryImpl(db)
}
