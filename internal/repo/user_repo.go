package repo

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"manageArticles_/internal/model"
)

type UserRepository interface {
	FindUser(email string, password string) (*model.User, error)
	CreateUser(user *model.User) error
	FindBtID(id int) (*model.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}
func (i *UserRepositoryImpl) FindBtID(id int) (*model.User, error) {
	user := new(model.User)
	err := i.db.Where("id=?", id).Find(user).Error
	return user, err
}
func (u *UserRepositoryImpl) CreateUser(user *model.User) error {

	err := u.db.Create(user).Error
	return err
}

func (u *UserRepositoryImpl) FindUser(email string, password string) (*model.User, error) {
	user := new(model.User)

	err := u.db.Where("email=? and password=?", email, password).Find(user).Error
	return user, err
}
