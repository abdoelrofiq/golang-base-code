package middlewares

import (
	model "golang-base-code/src/app/models"

	"gorm.io/gorm"
)

type UserMiddleware interface {
	Fetch() ([]model.User, error)
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

func (m *userMiddlewareBuilder) Fetch() ([]model.User, error) {
	m.Db.Joins("Profession").Preload("Books", "author != ?", "Random Book").Find(&users)

	return users, nil
}
