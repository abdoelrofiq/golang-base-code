package seeders

import (
	model "golang-base-code/src/app/models"

	"gorm.io/gorm"
)

type ProfessionSeederBuilder interface {
	InsertProfessionSeeder()
}

type professioneederConnection struct {
	conn *gorm.DB
}

func Professioneeder(conn *gorm.DB) ProfessionSeederBuilder {
	return &professioneederConnection{
		conn: conn,
	}
}

func (profession *professioneederConnection) InsertProfessionSeeder() {
	var professionTotal int64

	profession.conn.Model(&model.Profession{}).Count(&professionTotal)

	if professionTotal > 0 {
		return
	}

	var Profession = []model.Profession{
		{
			Id:   1,
			Name: "Programmer",
		},
		{
			Id:   2,
			Name: "Doctor",
		}}

	profession.conn.Create(Profession)
}
