package seeders

import "gorm.io/gorm"

func RunSeeder(connection *gorm.DB) {

	bookSeeder := BookSeeder(connection)
	bookSeeder.InsertBooksSeeder()

	userSeeder := UserSeeder(connection)
	userSeeder.InsertUsersSeeder()

}
