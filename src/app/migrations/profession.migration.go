package migrations

import (
	model "golang-base-code/src/app/models"

	"gorm.io/gorm"
)

type ProfessionMigrationBuilder interface {
	RunProfessionMigration()
}

func (profession *professionMigrationConnection) RunProfessionMigration() {
	CreateProfessionTable(profession)
}

type professionMigrationConnection struct {
	conn *gorm.DB
}

func ProfessionMigration(conn *gorm.DB) ProfessionMigrationBuilder {
	return &professionMigrationConnection{
		conn: conn,
	}
}

func CreateProfessionTable(profession *professionMigrationConnection) {
	isExists := profession.conn.Migrator().HasTable(model.Profession{})
	if isExists {
		return
	}

	profession.conn.AutoMigrate(model.Profession{})
}
