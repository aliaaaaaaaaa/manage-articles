package repo

import (
	"github.com/jinzhu/gorm"
	"manageArticles_/internal/model"
)

type ArticleRepository interface {
	CreateArticle(article model.Article, user model.User) (*model.Article, error)
	UpdateArticle(article model.Article) (*model.Article, error)
	DeleteArticle(id int) error
	ShowArticles(limit int, offset int) []model.Article
}
type ArticleRepositoryImpl struct {
	db *gorm.DB
}

func NewArticleRepositoryImpl(db *gorm.DB) *ArticleRepositoryImpl {
	return &ArticleRepositoryImpl{db: db}
}

func (a ArticleRepositoryImpl) CreateArticle(article model.Article, user model.User) (*model.Article, error) {
	article.UserID = user.Id
	err := a.db.Create(article).Error
	return &article, err
}

func (a ArticleRepositoryImpl) UpdateArticle(article model.Article) (*model.Article, error) {
	err := a.db.Save(&article).Error
	return &article, err
}

func (a ArticleRepositoryImpl) DeleteArticle(id int) error {
	article := model.Article{Model: gorm.Model{ID: uint(id)}}
	err := a.db.Delete(&article).Error
	return err
}

func (a ArticleRepositoryImpl) ShowArticles(limit int, offset int) []model.Article {
	var articles []model.Article
	a.db.Limit(limit).Offset(offset).Find(&articles)
	return articles
}
