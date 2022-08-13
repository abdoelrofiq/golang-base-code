package migrations

import (
	model "golang-base-code/src/app/models"

	"gorm.io/gorm"
)

type BookMigrationBuilder interface {
	CreateBooksTable()
	AddOwnerId()
}

type bookMigrationConnection struct {
	conn *gorm.DB
}

func BookMigration(conn *gorm.DB) BookMigrationBuilder {
	return &bookMigrationConnection{
		conn: conn,
	}
}

func (book *bookMigrationConnection) CreateBooksTable() {
	isExists := book.conn.Migrator().HasTable("books")
	if isExists {
		return
	}

	book.conn.AutoMigrate(model.Books{})
}

func (book *bookMigrationConnection) AddOwnerId() {
	type Books struct {
		OwnerId int `json:"owner_id" validate:"required"`
	}

	isExists := book.conn.Migrator().HasColumn(Books{}, "OwnerId")
	if isExists {
		return
	}

	book.conn.Migrator().AddColumn(Books{}, "OwnerId")
}
