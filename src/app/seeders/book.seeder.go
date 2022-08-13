package seeders

import (
	model "golang-base-code/src/app/models"

	"gorm.io/gorm"
)

type BooksSeederBuilder interface {
	InsertBooksSeeder()
}

type bookSeederConnection struct {
	conn *gorm.DB
}

func BookSeeder(conn *gorm.DB) BooksSeederBuilder {
	return &bookSeederConnection{
		conn: conn,
	}
}

func (book *bookSeederConnection) InsertBooksSeeder() {

	var books = []model.Books{
		{
			Title:       "Peekaboo Whats in the Snow",
			Price:       46800,
			Author:      "Tim Pelangi Indonesia",
			Publisher:   "PELANGI INDONESIA",
			PublishDate: "2022-03-01",
		},
		{
			Title:       "Pengantar Ilmu Tafsir",
			Price:       43250,
			Author:      "Drs. A. Fudlali",
			Publisher:   "Angkasa",
			PublishDate: "2005-01-01",
		},
		{
			Title:       "The Miracle Of Ikhlas",
			Price:       29325,
			Author:      "Anin DP",
			Publisher:   "Mueeza",
			PublishDate: "2021-02-01",
		}}

	book.conn.Create(books)
}
