package seeders

import (
	model "golang-base-code/src/app/models"

	"gorm.io/gorm"
)

type UsersSeederBuilder interface {
	InsertUsersSeeder()
}

type userSeederConnection struct {
	conn *gorm.DB
}

func UserSeeder(conn *gorm.DB) UsersSeederBuilder {
	return &userSeederConnection{
		conn: conn,
	}
}
func (user *userSeederConnection) InsertUsersSeeder() {

	var users = []model.User{
		{
			Name:  "Wahyu",
			Email: "wahyuagung26@gmail.com",
		},
		{
			Name:  "Agung",
			Email: "wahyu.agung@majoo.id",
		}}

	user.conn.Create(users)
}
