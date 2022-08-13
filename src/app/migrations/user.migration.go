package migrations

import (
	model "golang-base-code/src/app/models"

	"gorm.io/gorm"
)

type UserMigrationBuilder interface {
	CreateUsersTable()
}
type userMigrationConnection struct {
	conn *gorm.DB
}

func UserMigration(conn *gorm.DB) UserMigrationBuilder {
	return &userMigrationConnection{
		conn: conn,
	}
}

func (user *userMigrationConnection) CreateUsersTable() {
	isExists := user.conn.Migrator().HasTable("users")
	if isExists {
		return
	}

	user.conn.AutoMigrate(model.User{})
}
