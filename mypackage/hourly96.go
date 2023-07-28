package mypackage

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "strconv"
	// "time"
	"gorm.io/gorm"

)

type Hourlys struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Hourlys    []Hourly   `json:"list"  gorm:"foreignKey: ParentID"`
	CityHourly CityHourly `json:"city" gorm:"foreignKey: CityHourlyID"`
}
type CityHourly struct {
	CityHourlyID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name"`
	TimeZone int64  `json:"timezone"`
}
type Hourly struct {
	HourlyId             uint                  `gorm:"primaryKey;autoIncrement"`
	ParentID             uint    
	WeathersHourly       []WeathersHourly     `json:"weather" gorm:"foreignKey:WeatherHourlyID"`
	MainParametersHourly MainParametersHourly `json:"main" gorm:"foreignKey:MainHourlyID"`
	Time                 int64                `json:"dt"`
	// RainHourly           RainHourly           `json:"rain"`
}

type RainHourly struct {
	Rain float64 `json:"1h"`
}
type WeathersHourly struct {
	WeatherHourlyID     uint       `gorm:"primaryKey;autoIncrement"`
	Weather string `json:"description"`
}
type MainParametersHourly struct {
	MainHourlyID uint    `gorm:"primaryKey;autoIncrement"`
	Temperature float64 `json:"temp"`
	Feels       float64 `json:"feels_like"`
}

func HourlyData(cityName string, DB *gorm.DB) {
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

	var hourlys Hourlys

	json.Unmarshal(byteResult, &hourlys)
	// DB.Create(&hourlys)


}
