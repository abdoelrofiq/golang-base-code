package migrations

import (
	model "golang-base-code/src/app/models"

	"gorm.io/gorm"
)

type UserMigrationBuilder interface {
	RunUserMigration()
}

func (user *userMigrationConnection) RunUserMigration() {
	CreateUserTable(user)
	AddIsActive(user)
}

type userMigrationConnection struct {
	conn *gorm.DB
}

func UserMigration(conn *gorm.DB) UserMigrationBuilder {
	return &userMigrationConnection{
		conn: conn,
	}
}

func CreateUserTable(user *userMigrationConnection) {
	isExists := user.conn.Migrator().HasTable(&model.User{})
	if isExists {
		return
	}

	user.conn.AutoMigrate(model.User{})
}

func AddIsActive(user *userMigrationConnection) {
	isExists := user.conn.Migrator().HasColumn(&model.User{}, "IsActive")
	if isExists {
		return
	}

	user.conn.Migrator().AddColumn(&model.User{}, "IsActive")
}
