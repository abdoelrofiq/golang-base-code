package migrations

import (
	model "golang-base-code/src/app/models"

	"gorm.io/gorm"
)

type ProfessionMigrationBuilder interface {
	CreateProfessionTable()
}

type professionMigrationConnection struct {
	conn *gorm.DB
}

func ProfessionMigration(conn *gorm.DB) ProfessionMigrationBuilder {
	return &professionMigrationConnection{
		conn: conn,
	}
}

func (profession *professionMigrationConnection) CreateProfessionTable() {
	isExists := profession.conn.Migrator().HasTable(model.Profession{})
	if isExists {
		return
	}

	profession.conn.AutoMigrate(model.Profession{})
}
