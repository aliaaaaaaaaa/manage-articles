package test

import (
	"github.com/stretchr/testify/assert"
	"manageArticles_/internal/model"
	"testing"
)

func TestUserRepo(t *testing.T) {
	user := model.User{
		Name:     "test",
		Email:    "test@test.com",
		Password: "123456",
		IsActive: true,
	}
	_, err := userRepo.CreateUser(&user)
	assert.Empty(t, err)
	user.Password = "123456"
	findUser, err := userRepo.FindUser(user.Email, user.Password)
	assert.Empty(t, err)
	assert.Equal(t, findUser.Id, user.Id)
	assert.Equal(t, findUser.Name, user.Name)
}
