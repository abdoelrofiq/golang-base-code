package seeders

import "gorm.io/gorm"

func RunSeeder(connection *gorm.DB) {

	bookSeeder := BookSeeder(connection)
	bookSeeder.InsertBookSeeder()

	userSeeder := UserSeeder(connection)
	userSeeder.InsertUserSeeder()

	Professioneeder := Professioneeder(connection)
	Professioneeder.InsertProfessionSeeder()

}
