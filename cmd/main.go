package main

import (
	"encoding/json"
	"fmt"
	"github.com/adeemgoogle/gowork/mypackage"
	Serv "github.com/adeemgoogle/gowork/src/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

var DB *gorm.DB



func main() {
	// var cityName string = "Almaty"
	// mypackage.CurrentData(cityName)

	DB := Serv.Init()


	r := gin.Default()


	r.POST("/createCurrent/:name", func(c *gin.Context)  {
		cityName := c.Param("name")

	
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
        cityName := c.Param("name")

	
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

		resp2, err := http.Get("https://pro.openweathermap.org/data/2.5/forecast/daily?q=" + cityName + "&appid=51e51b22fb137270e2e89bd2bc7c4acc&units=metric")
		if err != nil {
			log.Fatal(err)
		}
		defer resp2.Body.Close()

		// Чтение ответа
		byteResult2, err := ioutil.ReadAll(resp2.Body)
		if err != nil {
			log.Fatal(err)
		}

		var dailys2 mypackage.Dailys

		json.Unmarshal(byteResult2, &dailys2)
	
		// DB.Create(&dailys2)
	
		c.JSON(http.StatusCreated, dailys)
	})
	r.POST("/createHourly/:name", func(c *gin.Context)  {
        cityName := c.Param("name")

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
        var city mypackage.City
		if err := DB.Table("cities").
		Where("name = ?", dailysName).First(&city).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Dailys not found2"})
			return
		}



		// Load the associated `Daily` and related data from the connected tables.
		if err := DB.Preload("City").
		Preload("Daily.MainParametersDaily").
		Preload("Daily.WeathersDaily").
		Preload("Daily").
		Where("id = ?", city.CityDailyID).First(&existingDailys).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Dailys not found2"})
			return
		}



        var updatedDailys mypackage.Dailys
		updatedDailys.ID = existingDailys.ID
		fmt.Println(updatedDailys.ID)
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
		for i:=0;i<7;i++ {
			updatedDailys.Daily[i].DailyID = existingDailys.Daily[i].DailyID
			updatedDailys.Daily[i].ParentDailyID = existingDailys.Daily[i].ParentDailyID
			updatedDailys.Daily[i].WeathersDaily[0].WeathersDailyID = existingDailys.Daily[i].WeathersDaily[0].WeathersDailyID
			updatedDailys.Daily[i].MainParametersDaily.MainParametersDailyID = existingDailys.Daily[i].MainParametersDaily.MainParametersDailyID
		}
        existingDailys = updatedDailys

        // Save the updated user to the database
        DB.Save(&existingDailys)


        // Return the updated user in the response
        c.JSON(http.StatusOK, existingDailys)
	})
	r.PUT("/allhourly/:name", func(c *gin.Context) {
        dailysName := c.Param("name")

        var existingHourlys mypackage.Hourlys
        var cityHourly mypackage.CityHourly
		if err := DB.Table("city_hourlies").
		Where("name = ?", dailysName).First(&cityHourly).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Hourlys not found1"})
			return
		}

		// Load the associated `Daily` and related data from the connected tables.
		if err := DB.Preload("CityHourly").
		Preload("Hourly.MainParametersHourly").
		Preload("Hourly.WeathersHourly").
		Preload("Hourly").
		Where("id = ?", cityHourly.CityHourlyID).First(&existingHourlys).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Dailys not found2"})
			return
		}

        var updatedHourlys mypackage.Hourlys
		updatedHourlys.ID = existingHourlys.ID
		fmt.Println(updatedHourlys.ID)
		var cityName string = "Almaty"

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

		json.Unmarshal(byteResult, &updatedHourlys)



        // Update the existing user with the new data
		for i:=0;i<7;i++ {
			updatedHourlys.Hourly[i].HourlyId = existingHourlys.Hourly[i].HourlyId
			updatedHourlys.Hourly[i].ParentID = existingHourlys.Hourly[i].ParentID
			updatedHourlys.Hourly[i].WeathersHourly[0].WeatherHourlyID = existingHourlys.Hourly[i].WeathersHourly[0].WeatherHourlyID
			updatedHourlys.Hourly[i].MainParametersHourly.MainHourlyID = existingHourlys.Hourly[i].MainParametersHourly.MainHourlyID
		}
        existingHourlys = updatedHourlys

        // Save the updated user to the database
        DB.Save(&existingHourlys)


        // Return the updated user in the response
        c.JSON(http.StatusOK, existingHourlys)
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



	r.Run() // listen and serve on 0.0.0.0:8080

	
}



