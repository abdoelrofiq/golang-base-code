package config

import (
	"golang-base-code/src/app/database"
	"golang-base-code/src/app/utilities"
)

func DatabaseConfig() database.DatabaseConfig {
	var databaseConfig = database.DatabaseConfig{
		DB_USERNAME: utilities.GetEnvValue("DB_USERNAME"),
		DB_PASSWORD: utilities.GetEnvValue("DB_PASSWORD"),
		DB_HOST:     utilities.GetEnvValue("DB_HOST"),
		DB_NAME:     utilities.GetEnvValue("DB_NAME"),
		DB_PORT:     utilities.GetEnvValue("DB_PORT"),
	}

	return databaseConfig
}
