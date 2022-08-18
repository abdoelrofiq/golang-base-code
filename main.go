package main

import (
	"golang-base-code/src/app/config"
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

func main() {
	e := echo.New()
	e.Validator = &RequestValidator{validator: validator.New()}

	connection, _ := database.ConnectDatabase(config.DatabaseConfig())

	routes.AppRoutes(e, connection)
	migrations.RunMigration(connection)
	seeders.RunSeeder(connection)

	e.Logger.Fatal(e.Start(":" + utilities.GetEnvValue("APP_PORT")))
}
