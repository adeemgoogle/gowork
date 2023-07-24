package main

import (
	"github.com/adeemgoogle/gowork/Serv"
	"github.com/adeemgoogle/gowork/mypackage"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"

)

var DB *gorm.DB




// func createHourlyHandler(c *gin.Context) {
//     var cityName string = "Almaty"
// 	mypackage.HourlyData(cityName, DB)

// 	c.JSON(200, gin.H{
// 		"message": "ping",
// 	})
// }

// func createDailyHandler(c *gin.Context) {
//     var cityName string = "Almaty"
// 	mypackage.DailyData(cityName, DB)

// 	c.JSON(200, gin.H{
// 		"message": "ping",
// 	})
// }

func main() {
	// var cityName string = "Almaty"
	// mypackage.CurrentData(cityName)

	DB := Serv.Init()

	//migrations
	//daily / 7 days  
	// DB.AutoMigrate(&mypackage.MainParametersDaily{}, &mypackage.WeathersDaily{}, &mypackage.City{}, &mypackage.Daily{},&mypackage.Dailys{})

	//hourly / 96 hours
	// DB.AutoMigrate(&mypackage.CityHourly{}, &mypackage.WeathersHourly{}, &mypackage.MainParametersHourly{}, &mypackage.Hourly{}, &mypackage.Hourlys{})

	// current
	// DB.AutoMigrate(&mypackage.CloudsCurrent{}, &mypackage.InfoSunCurrent{}, &mypackage.MainParametersCurrent{}, &mypackage.WindCurrent{}, &mypackage.WeathersCurrent{}, &mypackage.Current{})




	r := gin.Default()
	// r.POST("/createCurrent", createCurrentHandler)
	// r.POST("/api/createHourly", createHourlyHandler)
	// r.POST("/api/createDaily", createDailyHandler)
	r.GET("/pong", func(c *gin.Context) {
		// var cityName string = "Almaty"
		// mypackage.CurrentData(cityName, DB)


		city10 := []string{"Almaty", "London", "Nur-Sultan", "Moscow"}

		for i := 0; i<4;i++{
			mypackage.CurrentData(city10[i], DB)
		}

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/createCurrent", func(c *gin.Context)  {
		var cityName string = "Almaty"
		// mypackage.CurrentData(cityName, DB)
	
		// city10 := []string{"Almaty", "London", "Astana", "Moscow", "Madrid"}
	
		// for i := 0; i<5;i++{
		// 	mypackage.CurrentData(city10[i], DB)
		// }
	
		resp, err := http.Get("https://pro.openweathermap.org/data/2.5/weather?q=" + cityName + "&appid=51e51b22fb137270e2e89bd2bc7c4acc&units=metric")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
	
		// Чтение ответа
		byteResult, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	
		var Currents mypackage.Current
	
		json.Unmarshal(byteResult, &Currents)
	
		DB.Create(&Currents)
	
		// c.JSON(200, gin.H{
		// 	"message": "ping",
		// })
	
		c.JSON(http.StatusCreated, Currents)
	})
	r.PUT("/alldailys/:id", func(c *gin.Context) {
        dailysID := c.Param("id")

        var existingDailys mypackage.Dailys

        if err := DB.First(&existingDailys, dailysID).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Current not found"})
            return
        }

        var updatedDailys mypackage.Dailys
		var cityName string = "Almaty"

		resp, err := http.Get("https://pro.openweathermap.org/data/2.5/forecast/daily?q=" + cityName + "&appid=51e51b22fb137270e2e89bd2bc7c4acc&units=metric")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		// Чтение ответа
		byteResult, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		// var dailys Dailys

		json.Unmarshal(byteResult, &updatedDailys)
        

        // Update the existing user with the new data
        existingDailys  = updatedDailys

        // Save the updated user to the database
        DB.Save(&existingDailys)

        // Return the updated user in the response
        c.JSON(http.StatusOK, existingDailys)
    })

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



