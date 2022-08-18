package database

import (
	"golang-base-code/src/app/utilities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_NAME     string
	DB_PORT     string
}

func createMysqlDatabase(c DatabaseConfig) bool {
	dsn := c.DB_USERNAME + ":" + c.DB_PASSWORD + "@tcp(" + c.DB_HOST + ":" + c.DB_PORT + ")/" + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Printf("Can't connect to the database server")
	}

	createDatabaseQuery := "CREATE DATABASE `" + c.DB_NAME + "`;"
	db = db.Exec(createDatabaseQuery)
	if db.Error != nil {
		log.Printf("Can't create the database")
	}

	return true
}

func attemptMysqlConnection(c DatabaseConfig) (*gorm.DB, error) {
	dsn := c.DB_USERNAME + ":" + c.DB_PASSWORD + "@tcp(" + c.DB_HOST + ":" + c.DB_PORT + ")/" + c.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return db, err
	}

	return db, nil
}

func ConnectMysql(c DatabaseConfig) (*gorm.DB, error) {
	db, err := attemptMysqlConnection(c)
	if err != nil {
		migration := utilities.GetEnvValue("MIGRATION")

		if migration == "UP" {
			createMysqlDatabase(c)
			db, err = attemptMysqlConnection(c)
			if err != nil {
				return db, err
			}
		} else {
			return db, err
		}
	}

	return db, nil
}

func ConnectDatabase(c DatabaseConfig) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	databaseDriver := utilities.GetEnvValue("DB_DRIVER")

	if databaseDriver == "MYSQL" {
		db, err = ConnectMysql(c)
		if err != nil {
			log.Printf("Can't connect to the database")
		}

	}

	return db, nil
}
