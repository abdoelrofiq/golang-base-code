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
			Title:       "Peekaboo Whats in the Snow",
			Price:       46800,
			Author:      "Tim Pelangi Indonesia",
			Publisher:   "PELANGI INDONESIA",
			PublishDate: "2022-03-01",
			OwnerId:     1,
		},
		{
			Id:          2,
			Title:       "Pengantar Ilmu Tafsir",
			Price:       43250,
			Author:      "Drs. A. Fudlali",
			Publisher:   "Angkasa",
			PublishDate: "2005-01-01",
			OwnerId:     1,
		},
		{
			Id:          3,
			Title:       "The Miracle Of Ikhlas",
			Price:       29325,
			Author:      "Anin DP",
			Publisher:   "Mueeza",
			PublishDate: "2021-02-01",
			OwnerId:     2,
		}}

	book.conn.Create(books)
}
