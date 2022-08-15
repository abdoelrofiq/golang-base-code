package migrations

import (
	model "golang-base-code/src/app/models"

	"gorm.io/gorm"
)

type UserMigrationBuilder interface {
	CreateUserTable()
	AddIsActive()
}
type userMigrationConnection struct {
	conn *gorm.DB
}

func UserMigration(conn *gorm.DB) UserMigrationBuilder {
	return &userMigrationConnection{
		conn: conn,
	}
}

func (user *userMigrationConnection) CreateUserTable() {
	isExists := user.conn.Migrator().HasTable(&model.User{})
	if isExists {
		return
	}

	user.conn.AutoMigrate(model.User{})
}

func (user *userMigrationConnection) AddIsActive() {
	isExists := user.conn.Migrator().HasColumn(&model.User{}, "IsActive")
	if isExists {
		return
	}

	user.conn.Migrator().AddColumn(&model.User{}, "IsActive")
}
