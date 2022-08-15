package migrations

import "gorm.io/gorm"

func RunMigration(connection *gorm.DB) {

	bookMigration := BookMigration(connection)
	bookMigration.CreateBookTable()

	professionMigration := ProfessionMigration(connection)
	professionMigration.CreateProfessionTable()

	userMigration := UserMigration(connection)
	userMigration.CreateUserTable()
	userMigration.AddIsActive()

}
