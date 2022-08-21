package middlewares

import (
	"errors"
	"golang-base-code/src/app/core"
	model "golang-base-code/src/app/models"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

type UserMiddleware interface {
	Fetch(c echo.Context) ([]model.User, error)
}

type userMiddlewareBuilder struct {
	DB *gorm.DB
}

var users []model.User

func UserConnectionMw(connection *gorm.DB) UserMiddleware {
	return &userMiddlewareBuilder{
		DB: connection,
	}
}

func (conn *userMiddlewareBuilder) Fetch(c echo.Context) ([]model.User, error) {
	FQP, err := core.FQP(conn.DB, c)
	if err != nil {
		return users, errors.New(err.Error())
	}

	result := FQP.Joins("Profession").Preload("Books", "author != ?", "Random Book").Find(&users)
	if result.Error != nil {
		return users, errors.New(result.Error.Error())
	}

	return users, nil
}
