package user

import (
	"context"
	"errors"

	"github.com/daffashafwan/taskmajoo/business/users"
	"github.com/daffashafwan/taskmajoo/helpers/encrypt"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func CreateUserRepo(conn *gorm.DB) users.Repository {
	return &UserRepo{
		DB: conn,
	}
}

func (rep *UserRepo) Login(ctx context.Context, username string, password string) (users.Domain, error) {
	var user User
	result := rep.DB.Table("Users").Where("user_name = ?", username).First(&user).Error

	if result != nil {
		return users.Domain{}, result
	}
	if !(encrypt.Compare(password,user.Password)) {
		return users.Domain{}, errors.New("Password tidak cocok")
	}
	return user.ToDomain(), nil
}

