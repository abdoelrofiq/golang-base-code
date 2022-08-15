package migrations

import (
	model "golang-base-code/src/app/models"

	"gorm.io/gorm"
)

type BookMigrationBuilder interface {
	RunBookMigration()
}

func (book *bookMigrationConnection) RunBookMigration() {
	CreateBookTable(book)
}

type bookMigrationConnection struct {
	conn *gorm.DB
}

func BookMigration(conn *gorm.DB) BookMigrationBuilder {
	return &bookMigrationConnection{
		conn: conn,
	}
}

func CreateBookTable(book *bookMigrationConnection) {
	isExists := book.conn.Migrator().HasTable(model.Book{})
	if isExists {
		return
	}

	book.conn.AutoMigrate(model.Book{})
}
