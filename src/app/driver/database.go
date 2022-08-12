package driver

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDb struct {
	MYSQL_USERNAME string
	MYSQL_PASSWORD string
	MYSQL_HOST     string
	MYSQL_DB       string
	MYSQL_PORT     string
}

func ConnectMysql(c ConfigDb) (*gorm.DB, error) {
	dsn := c.MYSQL_USERNAME + ":" + c.MYSQL_PASSWORD + "@tcp(" + c.MYSQL_HOST + ":" + c.MYSQL_PORT + ")/" + c.MYSQL_DB + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect to database")
	}

	return db, nil
}
