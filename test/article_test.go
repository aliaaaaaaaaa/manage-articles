package test

import (
	"github.com/stretchr/testify/assert"
	"manageArticles_/internal/model"
	"testing"
)

func TestArticle(t *testing.T) {
	user := model.User{
		Id:   1,
		Name: "test",
	}
	article := model.Article{
		Title: "testMe",
		Tags:  nil,
	}
	createArticle, err := articleRepo.CreateArticle(article, user)
	assert.Empty(t, err)
	assert.Equal(t, createArticle.Author, user.Name)
	articles := articleRepo.ShowArticles(10, 0)
	assert.NotEmpty(t, articles)
}
