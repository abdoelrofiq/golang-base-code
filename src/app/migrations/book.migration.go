package migrations

import (
	model "golang-base-code/src/app/models"

	"gorm.io/gorm"
)

type BookMigrationBuilder interface {
	CreateBookTable()
}

type bookMigrationConnection struct {
	conn *gorm.DB
}

func BookMigration(conn *gorm.DB) BookMigrationBuilder {
	return &bookMigrationConnection{
		conn: conn,
	}
}

func (book *bookMigrationConnection) CreateBookTable() {
	isExists := book.conn.Migrator().HasTable(model.Book{})
	if isExists {
		return
	}

	book.conn.AutoMigrate(model.Book{})
}
