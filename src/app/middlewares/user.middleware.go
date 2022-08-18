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
	Db *gorm.DB
}

var users []model.User

func UserConnectionMw(connection *gorm.DB) UserMiddleware {
	return &userMiddlewareBuilder{
		Db: connection,
	}
}

func (m *userMiddlewareBuilder) Fetch(c echo.Context) ([]model.User, error) {
	var queryDB *gorm.DB

	filterQueryString, filterArgument, err := core.FQP(c)
	if err != nil {
		return users, errors.New(err.Error())
	}

	if len([]rune(filterQueryString)) == 0 && filterArgument == nil {
		queryDB = m.Db
	} else {
		queryDB = m.Db.Where(filterQueryString, filterArgument)
	}

	queryDB.Joins("Profession").Preload("Books", "author != ?", "Random Book").Find(&users)

	return users, nil
}
