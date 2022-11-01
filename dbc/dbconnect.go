package dbc

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

//var Db *gorm.DB

func Dbinit() *gorm.DB {
	dns := `postgresql://jzoxvjombiuzvjolvrtikkns@psql-mock-database-cloud:yrddycltuxgnukwshiiftbfi@psql-mock-database-cloud.postgres.database.azure.com:5432/ecom1666192780245qbwvcnhcdzktnoum`
	DB, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Println(err)
	} else {
		log.Println("connected")
	}
	return DB
}
