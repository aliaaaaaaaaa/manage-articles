package repo

import (
	"github.com/jinzhu/gorm"
	"manageArticles_/internal/model"
)

type ArticleRepository interface {
	CreateArticle(article *model.Article) error
	UpdateArticle(article *model.Article) error
	DeleteArticle(article model.Article) error
	ShowArticles(limit int, offset int, id int) []model.Article
	GetById(id int) (model.Article, error)
}
type ArticleRepositoryImpl struct {
	db *gorm.DB
}

func NewArticleRepositoryImpl(db *gorm.DB) *ArticleRepositoryImpl {
	return &ArticleRepositoryImpl{db: db}
}
func (a *ArticleRepositoryImpl) GetById(id int) (model.Article, error) {
	var article model.Article
	err := a.db.Where("id=?", id).Find(&article).Error
	return article, err
}
func (a *ArticleRepositoryImpl) CreateArticle(article *model.Article) error {
	err := a.db.Create(article).Error
	return err
}

func (a *ArticleRepositoryImpl) UpdateArticle(article *model.Article) error {
	err := a.db.Save(&article).Error
	return err
}

func (a *ArticleRepositoryImpl) DeleteArticle(article model.Article) error {
	err := a.db.Delete(&article).Error
	return err
}

func (a *ArticleRepositoryImpl) ShowArticles(limit int, offset int, id int) []model.Article {
	var articles []model.Article
	a.db.Limit(limit).Offset(offset).Where("user_id=?", id).Find(&articles)
	return articles
}
