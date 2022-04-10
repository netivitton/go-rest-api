package utils

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgreSQLconnection interface {
	Connection() gorm.DB
}

func Connection() (gormdb *gorm.DB, err error) {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal("cannot connect DB", err)
	}
	//dsn := "host=" + config.DB_HOST + " user=" + config.DB_USER + " password=" + config.DB_PASS + " dbname=" + config.DB_NAME + " port=9920 sslmode=disable TimeZone=Asia/Bangkok"
	dsn := config.DB_HOST_ALL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, err
}

func CloseDB(gormdb *gorm.DB) (err error) {
	sqlDB, err := gormdb.DB()
	sqlDB.Close()
	return err
}
