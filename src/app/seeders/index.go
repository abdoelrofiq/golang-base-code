package seeders

import (
	"golang-base-code/src/app/utilities"

	"gorm.io/gorm"
)

func RunSeeder(connection *gorm.DB) {
	migration := utilities.GetEnvValue("MIGRATION")

	if migration == "UP" {
		bookSeeder := BookSeeder(connection)
		bookSeeder.InsertBookSeeder()

		userSeeder := UserSeeder(connection)
		userSeeder.InsertUserSeeder()

		Professioneeder := Professioneeder(connection)
		Professioneeder.InsertProfessionSeeder()
	}
}
