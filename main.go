package main

import (
	"golang-base-code/src/app/database"
	"golang-base-code/src/app/migrations"
	"golang-base-code/src/app/routes"
	"golang-base-code/src/app/seeders"
	"golang-base-code/src/app/utilities"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type RequestValidator struct {
	validator *validator.Validate
}

func (cv *RequestValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

var config = database.ConfigDb{
	MYSQL_USERNAME: utilities.GetEnvValue("MYSQL_USERNAME"),
	MYSQL_PASSWORD: utilities.GetEnvValue("MYSQL_PASSWORD"),
	MYSQL_HOST:     utilities.GetEnvValue("MYSQL_HOST"),
	MYSQL_DB:       utilities.GetEnvValue("MYSQL_DB"),
	MYSQL_PORT:     utilities.GetEnvValue("MYSQL_PORT"),
}

func main() {
	e := echo.New()
	e.Validator = &RequestValidator{validator: validator.New()}

	connection, _ := database.ConnectMysql(config)

	routes.AppRoutes(e, connection)
	migrations.RunMigration(connection)
	seeders.RunSeeder(connection)

	// Start server
	e.Start(":4200")
}
