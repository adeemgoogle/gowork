package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Current struct {
	Id                    int                    `json:"id" gorm:"primaryKey"`
	NameCurrent           string                 `json:"name"`
	WeathersCurrent       []WeathersCurrent      `json:"weather" gorm:"foreignKey: WeatherID"`
	MainParametersCurrent *MainParametersCurrent `json:"main" gorm:"foreignKey: CurrentID"`
	WindCurrnet           WindCurrent            `json:"wind" gorm:"foreignKey: WindID"`
	InfoSunCurrent        *InfoSunCurrent        `json:"sys" gorm:"foreignKey:SunsetID"`
	TimeZone              int64                  `json:"timezone"`
	CloudsCurrent         *CloudsCurrent         `json:"clouds" gorm:"foreignKey: CloudsId"`
	Visibility            float64                `json:"visibility"`
}

type CloudsCurrent struct {
	CloudsId int `gorm:"primaryKey;autoIncrement"`
	Clouds   int `json:"all"`
}
type WeathersCurrent struct {
	WeatherID int `json:"id" gorm:"primaryKey"`
	// ID_weather int `json:"id_weather"`
	Main    string `json:"main"`
	Weather string `json:"description"`
	Icon    string `json:"icon"`
}
type MainParametersCurrent struct {
	CurrentID int     `gorm:"primaryKey;autoIncrement"`
	Current   float64 `json:"temp"`
	Feels     float64 `json:"feels_like"`
	Pressure  float64 `json:"pressure"`
	Humidity  float64 `json:"humidity"`
}
type WindCurrent struct {
	WindID    int     `gorm:"primaryKey;autoIncrement"`
	WindSpeed float64 `json:"speed"`
}
type InfoSunCurrent struct {
	SunsetID  int   `json:"id" gorm:"primaryKey"`
	RiseofSun int64 `json:"sunrise"`
	SetofSun  int64 `json:"sunset"`
}

func getCurrentApi(c *gin.Context) {

	cityID := c.Query("city_id") // Get city ID from the query parameter
	temperatureUnit := c.Query("temperature_unit")
	windSpeedUnit := c.Query("windSpeed_unit")
	pressureUnit := c.Query("pressure_unit")
	visionUnit := c.Query("vision_unit")
	// Connect to the PostgreSQL database (replace connection parameters with yours).
	dsn := "host=localhost user=postgres password=123 dbname=weather port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection error"})
		return
	}
	var Currents Current

	// Fetch the weather data based on the city ID
	if err := db.Preload("WeathersCurrent").Preload("MainParametersCurrent").Preload("WindCurrnet").Preload("InfoSunCurrent").Preload("CloudsCurrent").Where("id = ?", cityID).First(&Currents).Error; err != nil {
		fmt.Println("Error fetching data from the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data from the database"})
		return
	}

	switch temperatureUnit {
	case "celsius":
		Currents.MainParametersCurrent.Current = Currents.MainParametersCurrent.Current * 1
		Currents.MainParametersCurrent.Feels = Currents.MainParametersCurrent.Feels * 1
	case "kelvin":
		Currents.MainParametersCurrent.Current = Currents.MainParametersCurrent.Current + 273
		Currents.MainParametersCurrent.Feels = Currents.MainParametersCurrent.Feels + 273
	case "fahrenheit":
		Currents.MainParametersCurrent.Current = (Currents.MainParametersCurrent.Current * 1.8) + 32
		Currents.MainParametersCurrent.Feels = (Currents.MainParametersCurrent.Feels * 1.8) + 32

	}

	switch windSpeedUnit {
	case "Mps":
		Currents.WindCurrnet.WindSpeed = Currents.WindCurrnet.WindSpeed * 1
	case "kmh":
		Currents.WindCurrnet.WindSpeed = Currents.WindCurrnet.WindSpeed * 3.6
	case "mileh":
		Currents.WindCurrnet.WindSpeed = Currents.WindCurrnet.WindSpeed * 2.2369362912
	case "knots":
		Currents.WindCurrnet.WindSpeed = Currents.WindCurrnet.WindSpeed * 1.944
	}

	switch pressureUnit {
	case "pascal":
		Currents.MainParametersCurrent.Pressure = Currents.MainParametersCurrent.Pressure * 1
	case "mmHg":
		Currents.MainParametersCurrent.Pressure = Currents.MainParametersCurrent.Pressure * 0.00750062
	case "Mbar":
		Currents.MainParametersCurrent.Pressure = Currents.MainParametersCurrent.Pressure * 0.01
	}

	switch visionUnit {
	case "meter":
		Currents.Visibility = Currents.Visibility * 1
	case "mile":
		Currents.Visibility = Currents.Visibility * 0.000621371
	case "km":
		Currents.Visibility = Currents.Visibility * 0.001
	}
	c.JSON(http.StatusOK, Currents)
}

func RunAPI() {
	r := gin.Default()

	r.GET("/current/data/", getCurrentApi)

	r.Run("localhost:8080")
}