package main

import (
	"encoding/json"
	"fmt"
	"github.com/adeemgoogle/gowork/Fronts"
	"github.com/adeemgoogle/gowork/Serv"
	"github.com/adeemgoogle/gowork/mypackage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
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
	
		c.JSON(http.StatusCreated, Currents)
	})
	r.POST("/createDaily", func(c *gin.Context)  {
		var cityName string = "Almaty"
		// mypackage.CurrentData(cityName, DB)
	
		// city10 := []string{"Almaty", "London", "Astana", "Moscow", "Madrid"}
	
		// for i := 0; i<5;i++{
		// 	mypackage.CurrentData(city10[i], DB)
		// }
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

		var dailys mypackage.Dailys

		json.Unmarshal(byteResult, &dailys)
	
		DB.Create(&dailys)
	
		c.JSON(http.StatusCreated, dailys)
	})
	r.PUT("/alldailys/:name", func(c *gin.Context) {
        dailysName := c.Param("name")

        var existingDailys mypackage.Dailys

		var citi mypackage.City
		// db.Preload("Posts").Where("email = ?", userEmail).First(&user).Error;
		if err := DB.Table("cities").Where("name = ?", dailysName).First(&citi).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Dailys not found"})
            return
        }

        if err := DB.First(&existingDailys, citi.CityDailyID).Error; err != nil {
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
        // existingDailys  = updatedDailys

        // Save the updated user to the database
        // DB.Save(&existingDailys)
		DB.Model(&existingDailys).Updates(updatedDailys)

        // Return the updated user in the response
        c.JSON(http.StatusOK, existingDailys)
    })
	r.PUT("/allCurrent/:name", func(c *gin.Context) {
        currentName := c.Param("name")
		fmt.Println(currentName)

        var existingCurrent mypackage.Current

		// var citi mypackage.City


		// db.Preload("Posts").Where("email = ?", userEmail).First(&user).Error;

		if err := DB.Table("currents").Where("name_current  = ?", currentName).First(&existingCurrent).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Current not found"})
            return
        }

        // if err := DB.First(&existingDailys, citi.CityDailyID).Error; err != nil {
        //     c.JSON(http.StatusNotFound, gin.H{"error": "Current not found"})
        //     return
        // }

        var updatedCurrent mypackage.Current
		var cityName string = "Almaty"

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

		// var dailys Dailys

		json.Unmarshal(byteResult, &updatedCurrent)
        

        // Update the existing user with the new data
        // existingDailys  = updatedDailys

        // Save the updated user to the database
        // DB.Save(&existingDailys)
		// DB.Model(&existingCurrent).Updates(updatedCurrent)

        // Return the updated user in the response
        c.JSON(http.StatusOK, existingCurrent)
    })

	// r.GET("/daily", func(c *gin.Context) {
	// 	var cityName string = "Almaty"
	// 	mypackage.DailyData(cityName, DB)

	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	Fronts.GetCurrent()
	Fronts.GetDaily()
	Fronts.GetHourly()
	r.Run() // listen and serve on 0.0.0.0:8080

	// Serv.Close(DB)
}



