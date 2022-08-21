package main

import (
	"golang-base-code/src/app/config"
	"golang-base-code/src/app/database"
	"golang-base-code/src/app/migrations"
	"golang-base-code/src/app/routes"
	"golang-base-code/src/app/seeders"
	"golang-base-code/src/app/utilities"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	connection, _ := database.ConnectDatabase(config.DatabaseConfig())

	routes.AppRoutes(e, connection)
	migrations.RunMigration(connection)
	seeders.RunSeeder(connection)

	e.Logger.Fatal(e.Start(":" + utilities.GetEnvValue("APP_PORT")))
}
