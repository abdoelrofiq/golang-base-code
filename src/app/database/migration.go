package database

import "gorm.io/gorm"

func RunMigration(connection *gorm.DB) {
	// Import books seeder
	bookSeeder := BookMigration(connection)
	bookSeeder.ImportSeeder()

	// Import user seeder
	userSeeder := UserMigration(connection)
	userSeeder.ImportSeeder()
}
