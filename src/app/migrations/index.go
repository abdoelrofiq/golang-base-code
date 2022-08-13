package migrations

import "gorm.io/gorm"

func RunMigration(connection *gorm.DB) {

	bookMigration := BookMigration(connection)
	bookMigration.CreateBooksTable()
	bookMigration.AddOwnerId()

	userMigration := UserMigration(connection)
	userMigration.CreateUsersTable()

}
