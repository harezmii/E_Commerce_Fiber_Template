package database

import (
	"apirestecommerce/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB{
	dsn := "host=localhost user="+config.GetEnv("DATABASE_USER") +" password="+config.GetEnv("DATABASE_PASSWORD")+" dbname="+config.GetEnv("DATABASE_NAME")+" port="+config.GetEnv("DATABASE_PORT")+" sslmode=disable TimeZone="+config.GetEnv("DATABASE_TIMEZONE")
	db, dbError := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbError != nil {
		print("Database Error")
	}
	return db
}
