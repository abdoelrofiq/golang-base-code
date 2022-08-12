package migrations

import (
	model "golang-base-code/src/app/models"

	"gorm.io/gorm"
)

type BooksMigrationMigrator interface {
	ImportSeeder()
	AddOwnerId()
}

type bookMigrationBuilder struct {
	conn *gorm.DB
}

func BookMigration(conn *gorm.DB) BooksMigrationMigrator {
	return &bookMigrationBuilder{
		conn: conn,
	}
}

func (book *bookMigrationBuilder) ImportSeeder() {
	// Skip migration if books table already exist
	// and run migration if books table not exist
	isExists := book.conn.Migrator().HasTable("books")
	if isExists {
		return
	}

	// Prepare data for data dummy
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

	// Create table books and insert batch data dummy
	book.conn.AutoMigrate(model.Books{})
	book.conn.Create(books)
}

func (book *bookMigrationBuilder) AddOwnerId() {
	book.conn.Migrator().AddColumn(model.Books{}, "OwnerId")
}
