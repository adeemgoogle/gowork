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
	"fmt"
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
	r.POST("/createHourly", func(c *gin.Context)  {
		var cityName string = "Almaty"
		// mypackage.CurrentData(cityName, DB)
	
		// city10 := []string{"Almaty", "London", "Astana", "Moscow", "Madrid"}
	
		// for i := 0; i<5;i++{
		// 	mypackage.CurrentData(city10[i], DB)
		// }
	
		resp, err := http.Get("https://pro.openweathermap.org/data/2.5/forecast/hourly?q=" + cityName + "&appid=51e51b22fb137270e2e89bd2bc7c4acc&units=metric")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		// Чтение ответа
		byteResult, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var hourlys mypackage.Hourlys

		json.Unmarshal(byteResult, &hourlys)
	
		DB.Create(&hourlys)
	
		c.JSON(http.StatusCreated, hourlys)
	})
	r.PUT("/alldailys/:name", func(c *gin.Context) {
        dailysName := c.Param("name")

        var existingDailys mypackage.Dailys

		// var citi mypackage.City

		// db.Preload("Posts").Where("email = ?", userEmail).First(&user).Error;

		// if err := DB.Table("cities").Where("name = ?", dailysName).First(&citi).Error; err != nil {
        //     c.JSON(http.StatusNotFound, gin.H{"error": "Dailys not found"})
        //     return
        // }

        // if err := DB.First(&existingDailys, citi.CityDailyID).Error; err != nil {
        //     c.JSON(http.StatusNotFound, gin.H{"error": "Current not found"})
        //     return
        // }


		// Find the row in the `City` table where the `Name` column matches the given name.
		var city mypackage.City
		if err := DB.Where("name = ?", dailysName).First(&city).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Dailys not found1"})
			return
		}

		// Load the associated `Dailys` and related data from the connected tables.
		if err := DB.Preload("City").
		Preload("Daily").
		Where("id = ?", city.CityDailyID).First(&existingDailys).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Dailys not found2"})
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


		updatedDailys.ID = existingDailys.ID
		fmt.Println(updatedDailys.ID)
        

        // Update the existing user with the new data
        existingDailys  = updatedDailys

        // Save the updated user to the database
        // DB.Save(&existingDailys)
		// DB.Model(&existingDailys).Updates(updatedDailys)

		if err := DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&existingDailys).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the current"})
			return
		}

        // Return the updated user in the response
        c.JSON(http.StatusOK, existingDailys)
    })
	r.PUT("/allCurrent/:name", func(c *gin.Context) {
        currentName := c.Param("name")

        var existingCurrent mypackage.Current

		err := DB.Where("name_current = ?", currentName).Preload("WeathersCurrent").
		Preload("MainParametersCurrent").
		Preload("WindCurrnet").
		Preload("InfoSunCurrent").
		Preload("CloudsCurrent").
		First(&existingCurrent).Error

		if err != nil {
			// Handle the error
			c.JSON(http.StatusNotFound, gin.H{"error": "Current not found"})
            return
		}

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

		json.Unmarshal(byteResult, &updatedCurrent)


        // Update the existing user with the new data
        existingCurrent = updatedCurrent

        // Save the updated user to the database
		if err := DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&existingCurrent).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the current"})
			return
		}

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

	r.Run() // listen and serve on 0.0.0.0:8080

	// Serv.Close(DB)
}



