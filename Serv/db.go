package Serv

import (
	"fmt"
	"github.com/adeemgoogle/gowork/mypackage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:123@localhost:5432/weather"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")
	db.AutoMigrate(&mypackage.Current{})
	db.AutoMigrate(&mypackage.CloudsCurrent{})
	db.AutoMigrate(&mypackage.InfoSunCurrent{})
	db.AutoMigrate(&mypackage.MainParametersCurrent{})
	db.AutoMigrate(&mypackage.WindCurrnet{})
	db.AutoMigrate(&mypackage.WeathersCurrent{})
	
	return db
}
func Close(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		log.Println(err)
		return
	}
	err = dbSQL.Close()
	if err != nil {
		log.Println(err)
	}
}
