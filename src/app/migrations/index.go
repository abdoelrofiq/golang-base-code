package migrations

import (
	"golang-base-code/src/app/utilities"
	"log"

	"gorm.io/gorm"
)

func RunMigration(connection *gorm.DB) {
	migration := utilities.GetEnvValue("MIGRATION")

	if migration == "UP" {
		bookMigration := BookMigration(connection)
		bookMigration.RunBookMigration()

		professionMigration := ProfessionMigration(connection)
		professionMigration.RunProfessionMigration()

		userMigration := UserMigration(connection)
		userMigration.RunUserMigration()
	} else if migration == "DOWN" {
		dropDatabaseQuery := "DROP DATABASE `" + utilities.GetEnvValue("DB_NAME") + "`;"
		connection = connection.Exec(dropDatabaseQuery)
		if connection.Error != nil {
			log.Printf("Can't drop the database")
		}
	}
}
