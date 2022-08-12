package migrations

import "gorm.io/gorm"

func RunMigration(connection *gorm.DB) {
	// Import books seeder
	bookSeeder := BookMigration(connection)
	bookSeeder.ImportSeeder()
	bookSeeder.AddOwnerId()

	// Import user seeder
	userSeeder := UserMigration(connection)
	userSeeder.ImportSeeder()

}
