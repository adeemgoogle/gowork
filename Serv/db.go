package Serv

import (
	"fmt"
	"github.com/adeemgoogle/gowork/pkg"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:123@localhost:5432/weather"

	//dsn := fmt.Sprintf("host=localhost user=postgres password=clasypro04 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Almaty")


	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")

	// current
	// db.AutoMigrate(&mypackage.CloudsCurrent{})
	// db.AutoMigrate(&mypackage.InfoSunCurrent{})
	// db.AutoMigrate(&mypackage.MainParametersCurrent{})
	// db.AutoMigrate(&mypackage.WindCurrent{})
	// db.AutoMigrate(&mypackage.WeathersCurrent{})
	// db.AutoMigrate(&mypackage.Current{})


	//hourly / 96 hours
	db.AutoMigrate(&mypackage.CityHourly{})
	db.AutoMigrate(&mypackage.WeathersHourly{})
	db.AutoMigrate(&mypackage.MainParametersHourly{})
	db.AutoMigrate(&mypackage.Hourly{})
	db.AutoMigrate(&mypackage.Hourlys{})


	//daily / 7 days
	db.AutoMigrate(&mypackage.MainParametersDaily{})
	db.AutoMigrate(&mypackage.WeathersDaily{})
	db.AutoMigrate(&mypackage.City{})
	db.AutoMigrate(&mypackage.Daily{})
	db.AutoMigrate(&mypackage.Dailys{})

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
