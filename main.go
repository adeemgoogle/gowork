package main

import (
	"github.com/adeemgoogle/gowork/Serv"
	"github.com/adeemgoogle/gowork/mypackage"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

var DB *gorm.DB


func createCurrentHandler(c *gin.Context) {
    var cityName string = "Almaty"
	mypackage.CurrentData(cityName, DB)

	c.JSON(200, gin.H{
		"message": "ping",
	})
}

func createHourlyHandler(c *gin.Context) {
    var cityName string = "Almaty"
	mypackage.HourlyData(cityName, DB)

	c.JSON(200, gin.H{
		"message": "ping",
	})
}

func createDailyHandler(c *gin.Context) {
    var cityName string = "Almaty"
	mypackage.DailyData(cityName, DB)

	c.JSON(200, gin.H{
		"message": "ping",
	})
}

func main() {
	// var cityName string = "Almaty"
	// mypackage.CurrentData(cityName)

	DB := Serv.Init()

	//migrations
	//daily / 7 days  
	DB.AutoMigrate(&mypackage.MainParametersDaily{}, &mypackage.WeathersDaily{}, &mypackage.City{}, &mypackage.Daily{},&mypackage.Dailys{})

	//hourly / 96 hours
	DB.AutoMigrate(&mypackage.CityHourly{}, &mypackage.WeathersHourly{}, &mypackage.MainParametersHourly{}, &mypackage.Hourly{}, &mypackage.Hourlys{})

	// current
	DB.AutoMigrate(&mypackage.CloudsCurrent{}, &mypackage.InfoSunCurrent{}, &mypackage.MainParametersCurrent{}, &mypackage.WindCurrent{}, &mypackage.WeathersCurrent{}, &mypackage.Current{})





	r := gin.Default()
	r.POST("/api/createCurrent", createCurrentHandler)
	r.POST("/api/createHourly", createHourlyHandler)
	r.POST("/api/createDaily", createDailyHandler)
	// r.GET("/pong", func(c *gin.Context) {
	// 	var cityName string = "Almaty"
	// 	mypackage.HourlyData(cityName, DB)

	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.GET("/daily", func(c *gin.Context) {
	// 	var cityName string = "Almaty"
	// 	mypackage.DailyData(cityName, DB)

	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	r.Run() // listen and serve on 0.0.0.0:8080

	// Serv.Close(DB)
}



