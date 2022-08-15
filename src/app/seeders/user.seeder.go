package seeders

import (
	model "golang-base-code/src/app/models"

	"gorm.io/gorm"
)

type UserSeederBuilder interface {
	InsertUserSeeder()
}

type userSeederConnection struct {
	conn *gorm.DB
}

func UserSeeder(conn *gorm.DB) UserSeederBuilder {
	return &userSeederConnection{
		conn: conn,
	}
}

func (user *userSeederConnection) InsertUserSeeder() {
	var userTotal int64

	user.conn.Model(&model.Profession{}).Count(&userTotal)

	if userTotal > 0 {
		return
	}

	var users = []model.User{
		{
			Id:           1,
			Name:         "Wahyu",
			Email:        "wahyuagung26@gmail.com",
			ProfessionId: 1,
		},
		{
			Id:           2,
			Name:         "Agung",
			Email:        "wahyu.agung@majoo.id",
			ProfessionId: 1,
		}}

	user.conn.Create(users)
}
