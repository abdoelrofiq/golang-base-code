package domain

import (
	model "golang-base-code/src/app/models"
	repository "golang-base-code/src/http/repository/users"

	"gorm.io/gorm"
)

type mysqlUserDomain struct {
	Db *gorm.DB
}

var users []model.User

func MysqlUserDomain(connection *gorm.DB) repository.UserRepo {
	return &mysqlUserDomain{
		Db: connection,
	}
}

func (m *mysqlUserDomain) Fetch() ([]model.User, error) {
	m.Db.Raw("SELECT * FROM users").Scan(&users)

	return users, nil
}
