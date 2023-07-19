package Serv

import (
	"fmt"
	"github.com/adeemgoogle/gowork/mypackage"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init() *gorm.DB {
	// dbURL := "postgres://postgres:123@localhost:5432/weather"

	dsn := fmt.Sprintf("host=localhost user=postgres password=clasypro04 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Almaty")


	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")

	db.AutoMigrate(&mypackage.CloudsCurrent{})
	db.AutoMigrate(&mypackage.InfoSunCurrent{})
	db.AutoMigrate(&mypackage.MainParametersCurrent{})
	db.AutoMigrate(&mypackage.WindCurrnet{})
	db.AutoMigrate(&mypackage.WeathersCurrent{})
	db.AutoMigrate(&mypackage.Current{})


	//hourly / 96 hours
	db.AutoMigrate(&mypackage.Hourlys{})
	
	
	

	
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
