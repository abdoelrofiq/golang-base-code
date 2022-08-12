package database

import (
	model "golang-base-code/src/app/models"
	repository "golang-base-code/src/http/repository/users"

	"gorm.io/gorm"
)

type userMigration struct {
	conn *gorm.DB
}

func UserMigration(conn *gorm.DB) repository.UserMigrationRepo {
	return &userMigration{
		conn: conn,
	}
}

func (user *userMigration) ImportSeeder() {
	// Skip migration if users table already exist
	// and run migration if users table not exist
	isExists := user.conn.Migrator().HasTable("users")
	if isExists {
		return
	}

	// Prepare data for data dummy
	var users = []model.User{
		{
			Name:  "Wahyu",
			Email: "wahyuagung26@gmail.com",
		},
		{
			Name:  "Agung",
			Email: "wahyu.agung@majoo.id",
		}}

	// Create table users and insert batch data dummy
	user.conn.AutoMigrate(model.User{})
	user.conn.Create(users)
}
