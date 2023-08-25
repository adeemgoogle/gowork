package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	Serv "github.com/adeemgoogle/gowork/src/database"
	modell "github.com/adeemgoogle/gowork/src/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func PostCurrentData(r *gin.Engine) {
	DB = Serv.Init()
	r.POST("/createCurrent/:name", func(c *gin.Context) {
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

		var Currents modell.Current

		json.Unmarshal(byteResult, &Currents)

		DB.Create(&Currents)

		c.JSON(http.StatusCreated, Currents)
	})
}

func PutCurrentData(r *gin.Engine) {
	DB := Serv.Init()
	r.PUT("/allCurrent/:name", func(c *gin.Context) {
		currentName := c.Param("name")

		var existingCurrent modell.Current

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

		var updatedCurrent modell.Current
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
}

func PostHourlyData(r *gin.Engine) {
	DB = Serv.Init()
	r.POST("/createHourly/:name", func(c *gin.Context) {
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

		var hourlys modell.Hourlys

		json.Unmarshal(byteResult, &hourlys)

		DB.Create(&hourlys)

		c.JSON(http.StatusCreated, hourlys)
	})
}

func PutHourlyData(r *gin.Engine) {
	DB := Serv.Init()
	r.PUT("/allhourly/:name", func(c *gin.Context) {
		dailysName := c.Param("name")

		var existingHourlys modell.Hourlys
		var cityHourly modell.CityHourly
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

		var updatedHourlys modell.Hourlys
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
		for i := 0; i < 7; i++ {
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
}

func PostDailyData(r *gin.Engine) {
	DB := Serv.Init()
	r.POST("/createDaily/:name", func(c *gin.Context) {
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

		var dailys modell.Dailys

		json.Unmarshal(byteResult, &dailys)

		DB.Create(&dailys)

		c.JSON(http.StatusCreated, dailys)
	})
}

func PutDailyData(r *gin.Engine) {
	DB := Serv.Init()
	r.PUT("/alldailys/:name", func(c *gin.Context) {
		dailysName := c.Param("name")

		var existingDailys modell.Dailys
		var city modell.City
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

		var updatedDailys modell.Dailys
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
		for i := 0; i < 7; i++ {
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
}
