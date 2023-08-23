package Serv

import (
	"fmt"
<<<<<<< Updated upstream
	"github.com/adeemgoogle/gowork/mypackage"
=======
	"github.com/adeemgoogle/gowork/pkg/current"
	"github.com/adeemgoogle/gowork/pkg/hourly"

	//"github.com/adeemgoogle/gowork/pkg"
>>>>>>> Stashed changes
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Init() *gorm.DB {
	// dbURL := "postgres://postgres:123@localhost:5432/weather"
<<<<<<< Updated upstream

	dsn := fmt.Sprintf("host=localhost user=postgres password=clasypro04 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Almaty")
=======
	var dbhost,dbport,dbname,dbuser,dbpass string;
	dbhost = getEnv("DATABASE_HOST", "localhost")
	dbport = getEnv("DATABASE_PORT", "5432")
	dbname = getEnv("DATABASE_NAME", "weather")
	dbuser = getEnv("DATABASE_USERNAME", "postgres")
	dbpass = getEnv("DATABASE_PASSWORD", "123")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Almaty", dbhost, dbuser, dbpass, dbname, dbport)
>>>>>>> Stashed changes


	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")

	// current
<<<<<<< HEAD:src/database/db.go
<<<<<<< Updated upstream
	db.AutoMigrate(&mypackage.CloudsCurrent{})
	db.AutoMigrate(&mypackage.InfoSunCurrent{})
	db.AutoMigrate(&mypackage.MainParametersCurrent{})
	db.AutoMigrate(&mypackage.WindCurrent{})
	db.AutoMigrate(&mypackage.WeathersCurrent{})
	db.AutoMigrate(&mypackage.Current{})
=======
	db.AutoMigrate(&current.CloudsCurrent{})
	db.AutoMigrate(&current.InfoSunCurrent{})
	db.AutoMigrate(&current.MainParametersCurrent{})
	db.AutoMigrate(&current.WindCurrent{})
	db.AutoMigrate(&current.WeathersCurrent{})
	db.AutoMigrate(&current.WeathersCurrent{})
>>>>>>> Stashed changes
=======
	// db.AutoMigrate(&mypackage.CloudsCurrent{})
	// db.AutoMigrate(&mypackage.InfoSunCurrent{})
	// db.AutoMigrate(&mypackage.MainParametersCurrent{})
	// db.AutoMigrate(&mypackage.WindCurrent{})
	// db.AutoMigrate(&mypackage.WeathersCurrent{})
	// db.AutoMigrate(&mypackage.Current{})
>>>>>>> 783bd24564fe6e5f6f3f265b27729599ba7acac5:Serv/db.go


	//hourly / 96 hours
	db.AutoMigrate(&mypackage.CityHourly{})
	db.AutoMigrate(&mypackage.WeathersHourly{})
	db.AutoMigrate(&mypackage.MainParametersHourly{})
	db.AutoMigrate(&mypackage.Hourly{})
	db.AutoMigrate(&mypackage.Hourlys{})


	//daily / 7 days
	// db.AutoMigrate(&mypackage.MainParametersDaily{})
	// db.AutoMigrate(&mypackage.WeathersDaily{})
	// db.AutoMigrate(&mypackage.City{})
	// db.AutoMigrate(&mypackage.Daily{})
	// db.AutoMigrate(&mypackage.Dailys{})
<<<<<<< Updated upstream
=======






>>>>>>> Stashed changes



	
	

	
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