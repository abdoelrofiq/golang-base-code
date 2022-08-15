package database

import (
	"golang-base-code/src/app/utilities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ConfigDb struct {
	MYSQL_USERNAME string
	MYSQL_PASSWORD string
	MYSQL_HOST     string
	MYSQL_DB       string
	MYSQL_PORT     string
}

func createMysqlDatabase(c ConfigDb) bool {
	dsn := c.MYSQL_USERNAME + ":" + c.MYSQL_PASSWORD + "@tcp(" + c.MYSQL_HOST + ":" + c.MYSQL_PORT + ")/" + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Printf("Can't connect to the database server")
	}

	createDatabaseQuery := "CREATE DATABASE `" + c.MYSQL_DB + "`;"
	db = db.Exec(createDatabaseQuery)
	if db.Error != nil {
		log.Printf("Can't create the database")
	}

	return true
}

func attemptMysqlConnection(c ConfigDb) (*gorm.DB, error) {
	dsn := c.MYSQL_USERNAME + ":" + c.MYSQL_PASSWORD + "@tcp(" + c.MYSQL_HOST + ":" + c.MYSQL_PORT + ")/" + c.MYSQL_DB + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return db, err
	}

	return db, nil
}

func ConnectMysql(c ConfigDb) (*gorm.DB, error) {
	db, err := attemptMysqlConnection(c)
	if err != nil {
		migration := utilities.GetEnvValue("MIGRATION")

		if migration == "UP" {
			createMysqlDatabase(c)
			db, err = attemptMysqlConnection(c)
			if err != nil {
				log.Printf("Can't connect to the database")
			}
		} else {
			log.Printf("Can't connect to the database")
		}
	}

	return db, nil
}
