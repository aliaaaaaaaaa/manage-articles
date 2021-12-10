package service

import (
	"manageArticles_/internal/errors"
	"manageArticles_/internal/model"
	"manageArticles_/internal/repo"
)

type ArticleService struct {
	articleRepo repo.ArticleRepository
	userRepo    repo.UserRepository
}

func NewArticleService(articleRepo repo.ArticleRepository, userRepo repo.UserRepository) *ArticleService {
	return &ArticleService{articleRepo: articleRepo, userRepo: userRepo}
}
func (a *ArticleService) GetAllArticle(id int, limit int, offset int) []model.Article {
	articles := a.articleRepo.ShowArticles(limit, offset, id)
	return articles
}
func (a *ArticleService) CreateArticle(article *model.Article, userId int) error {
	user, err := a.userRepo.FindBtID(userId)
	if err != nil {
		return err
	}
	article.UserID = user.Id
	article.Author = user.Name
	err = a.articleRepo.CreateArticle(article)
	return err
}
func (a *ArticleService) DeleteArticle(id int, userId int) error {
	article, err := a.articleRepo.GetById(id)
	if err != nil {
		return err
	}
	if article.UserID != userId {
		return errors.NotPermit
	}
	err = a.articleRepo.DeleteArticle(article)
	return err
}
func (a *ArticleService) UpdateArticle(article *model.Article, userId int) error {
	getArticle, err := a.articleRepo.GetById(article.Id)
	if err != nil {
		return err
	}
	if getArticle.UserID != userId {
		return errors.NotPermit
	}
	user, err := a.userRepo.FindBtID(userId)
	if err != nil {
		return err
	}
	article.Author = user.Name
	article.UserID = user.Id
	err = a.articleRepo.UpdateArticle(article)
	return err
}
