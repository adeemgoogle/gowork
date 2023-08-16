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

// 	var currents Current

func GetCurrent() {
	// Подключение к базе данных PostgreSQL (замените параметры подключения на ваши).
	dsn := "host=localhost user=postgres password=123 dbname=weather port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return
	}

	// Получение данных из базы данных с объединением таблиц и предварительной загрузкой связанных данных.
	if err := db.Preload("WeathersCurrent").Preload("MainParametersCurrent").Preload("WindCurrnet").Preload("InfoSunCurrent").Preload("CloudsCurrent").Find(&Currents).Error; err != nil {
		fmt.Println("Ошибка при получении данных из базы данных:", err)
		return
	}

}

var Currents Current

func getCurrentApi(c *gin.Context) {

	GetCurrent()
	paramNameTemp := c.Param("paramNameTemp")
	paramNameWind := c.Param("paramNameWind")
	paramNamePress := c.Param("paramNamePress")
	paramNameVision := c.Param("paramNameVision")
	switch paramNameTemp {
	case "c":
		Currents.MainParametersCurrent.Current = Currents.MainParametersCurrent.Current * 1
		Currents.MainParametersCurrent.Feels = Currents.MainParametersCurrent.Feels * 1
	case "k":
		Currents.MainParametersCurrent.Current = Currents.MainParametersCurrent.Current + 273
		Currents.MainParametersCurrent.Feels = Currents.MainParametersCurrent.Feels + 273
	case "f":
		Currents.MainParametersCurrent.Current = (Currents.MainParametersCurrent.Current * 1.8) + 32
		Currents.MainParametersCurrent.Feels = (Currents.MainParametersCurrent.Feels * 1.8) + 32

	}

	switch paramNameWind {
	case "mm":
		Currents.WindCurrnet.WindSpeed = Currents.WindCurrnet.WindSpeed * 1
	case "kmh":
		Currents.WindCurrnet.WindSpeed = Currents.WindCurrnet.WindSpeed * 3.6
	case "mileh":
		Currents.WindCurrnet.WindSpeed = Currents.WindCurrnet.WindSpeed * 2.2369362912
	case "knots":
		Currents.WindCurrnet.WindSpeed = Currents.WindCurrnet.WindSpeed * 1.944

	}

	switch paramNamePress {
	case "Pa":
		Currents.MainParametersCurrent.Pressure = Currents.MainParametersCurrent.Pressure * 1
	case "mmHg":
		Currents.MainParametersCurrent.Pressure = Currents.MainParametersCurrent.Pressure * 0.00750062
	case "Mbar":
		Currents.MainParametersCurrent.Pressure = Currents.MainParametersCurrent.Pressure * 0.01

	}

	switch paramNameVision {
	case "meter":
		Currents.Visibility = Currents.Visibility * 1
	case "mile":
		Currents.Visibility = Currents.Visibility * 0.000621371
	case "km":
		Currents.Visibility = Currents.Visibility * 0.001
	}
	c.JSON(http.StatusOK, Currents)

}

// Добавляем новый роут, который будет получать данные о конкретном городе по его ID
func getCityData(c *gin.Context) {
	// Получаем параметр "cityID" из URL
	cityName := c.Param("cityName")

	// Подключение к базе данных PostgreSQL (замените параметры подключения на ваши).
	dsn := "host=localhost user=postgres password=123 dbname=weather port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при подключении к базе данных"})
		return
	}

	var city Current

	// Ищем данные о городе в базе данных по его ID
	if err := db.Preload("WeathersCurrent").Preload("MainParametersCurrent").Preload("WindCurrnet").Preload("InfoSunCurrent").Preload("CloudsCurrent").First(&city, cityName).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Город не найден"})
		return
	}

	// Возвращаем данные о городе в ответе
	c.JSON(http.StatusOK, city)
}

// Запускаем API
func RunApi() {
	r := gin.Default()

	// GET route для получения данных о конкретном городе по его ID
	r.GET("/api/cities/:cityID", getCityData)

	// GET route с параметрами "paramName" для получения данных о текущей погоде
	r.GET("/api/data/:paramNameTemp/:paramNameWind/:paramNamePress/:paramNameVision", getCurrentApi)

	r.Run("localhost:8080")
}

//33
//func RunApi() {
//	r := gin.Default()
//
//	// GET route with a parameter "paramName"
//	r.GET("/api/data/:paramNameTemp/:paramNameWind/:paramNamePress/:paramNameVision", getCurrentApi)
//	r.Run("localhost:8080")
//}
