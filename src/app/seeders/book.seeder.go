package seeders

import (
	model "golang-base-code/src/app/models"

	"gorm.io/gorm"
)

type BookSeederBuilder interface {
	InsertBookSeeder()
}

type bookSeederConnection struct {
	conn *gorm.DB
}

func BookSeeder(conn *gorm.DB) BookSeederBuilder {
	return &bookSeederConnection{
		conn: conn,
	}
}

func (book *bookSeederConnection) InsertBookSeeder() {
	var bookTotal int64

	book.conn.Model(&model.Book{}).Count(&bookTotal)

	if bookTotal > 0 {
		return
	}

	var books = []model.Book{
		{
			Id:          1,
			Title:       "Aku",
			Price:       50000,
			Author:      "Unkown",
			Publisher:   "Unkown",
			PublishDate: "2022-08-15",
			OwnerId:     1,
		}}

	book.conn.Create(books)
}
