package models

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		"pg-william-testing-williams-project-super-indo.l.aivencloud.com",
		"13608",
		"avnadmin",
		"AVNS_z0l3yS9e-uF014T_moO",
		"defaultdb",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to DB", err)
	}

	log.Println("connected")
	DB = db
}

func DBMigrate() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Category{})
	DB.AutoMigrate(&Product{})
	DB.Create(&Categories)
	DB.Create(&Products)
}
