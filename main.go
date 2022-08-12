package main

import (
	"golang-base-code/src/app/database"
	"golang-base-code/src/app/migrations"
	"golang-base-code/src/app/routes"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	"log"
	"os"

	"github.com/joho/godotenv"
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

func getEnvValue(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var config = database.ConfigDb{
	MYSQL_USERNAME: getEnvValue("MYSQL_USERNAME"),
	MYSQL_PASSWORD: getEnvValue("MYSQL_PASSWORD"),
	MYSQL_HOST:     getEnvValue("MYSQL_HOST"),
	MYSQL_DB:       getEnvValue("MYSQL_DB"),
	MYSQL_PORT:     getEnvValue("MYSQL_PORT"),
}

func main() {
	e := echo.New()
	e.Validator = &RequestValidator{validator: validator.New()}

	connection, _ := database.ConnectMysql(config)

	routes.AppRoutes(e, connection)
	migrations.RunMigration(connection)

	// Start server
	e.Start(":4200")
}
