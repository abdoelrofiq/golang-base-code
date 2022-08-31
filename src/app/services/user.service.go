package services

import (
	"errors"
	"golang-base-code/src/app/core"
	model "golang-base-code/src/app/models"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

type UserService interface {
	Fetch(c echo.Context) ([]model.User, error)
}

type userServiceBuilder struct {
	DB *gorm.DB
}

var users []model.User

func UserConnectionMw(connection *gorm.DB) UserService {
	return &userServiceBuilder{
		DB: connection,
	}
}

func (conn *userServiceBuilder) Fetch(c echo.Context) ([]model.User, error) {
	currentUser := c.Get("currentUser").(model.User)

	FQP, err := core.FQP(conn.DB, c)
	if err != nil {
		return users, errors.New(err.Error())
	}

	result := FQP.Where("users.id != ?", currentUser.Id).Joins("Profession").Preload("Books", "author != ?", "Random Book").Find(&users)
	if result.Error != nil {
		return users, errors.New(result.Error.Error())
	}

	return users, nil
}
