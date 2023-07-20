package mypackage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"gorm.io/gorm"

)

type Current struct {
	Id             int                   `json:"id" gorm:"primaryKey"`
	NameCurrent           string                `json:"name"`
	WeathersCurrent       []WeathersCurrent     `json:"weather" gorm:"foreignKey: WeatherID"`
	MainParametersCurrent MainParametersCurrent `json:"main" gorm:"foreignKey: CurrentID"`
	WindCurrnet           WindCurrent           `json:"wind" gorm:"foreignKey: WindID"`
	InfoSunCurrent        InfoSunCurrent        `json:"sys" gorm:"foreignKey:SunsetID"`
	TimeZone              int64                 `json:"timezone"`
	CloudsCurrent         CloudsCurrent         `json:"clouds" gorm:"foreignKey: CloudsId"`
	Visibility            int                   `json:"visibility"`
}

type CloudsCurrent struct {
	CloudsId int `gorm:"primaryKey;autoIncrement"`
	Clouds int `json:"all"`
}
type WeathersCurrent struct {
	WeatherID int `json:"id" gorm:"primaryKey"`
	// ID_weather int `json:"id_weather"`
	Main string `json:"main"`
	Weather string `json:"description"`
	Icon string `json:"icon"`
}
type MainParametersCurrent struct {
	CurrentID int `gorm:"primaryKey;autoIncrement"`
	Current  float64 `json:"temp"`
	Feels    float64 `json:"feels_like"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
}
type WindCurrent struct {
	WindID int `gorm:"primaryKey;autoIncrement"`
	WindSpeed int `json:"speed"`
}
type InfoSunCurrent struct {
	SunsetID int `json:"id" gorm:"primaryKey"`
	RiseofSun int64 `json:"sunrise"`
	SetofSun  int64 `json:"sunset"`
}

func CurrentData(cityName string, DB *gorm.DB) {
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

	var Currents Current

	json.Unmarshal(byteResult, &Currents)
	


	fmt.Println("ID: " + strconv.Itoa(Currents.Id))
	fmt.Println("Name: " + Currents.NameCurrent)
	fmt.Println("Weather Condition: " + Currents.WeathersCurrent[0].Weather)
	fmt.Println("Current Temperature: " + strconv.FormatFloat(Currents.MainParametersCurrent.Current, 'f', -2, 64))
	fmt.Println("Feels Like: " + strconv.FormatFloat(Currents.MainParametersCurrent.Feels, 'f', 0, 64))
	fmt.Println("Pressure : " + strconv.Itoa(Currents.MainParametersCurrent.Pressure))
	fmt.Println("Wind Speed: " + strconv.Itoa(Currents.WindCurrnet.WindSpeed))
	newSunRise := time.Unix(Currents.InfoSunCurrent.RiseofSun+Currents.TimeZone, 0).UTC()
	fmt.Println(newSunRise.Format("15:04:05"))
	newSunSet := time.Unix(Currents.InfoSunCurrent.SetofSun+Currents.TimeZone, 0).UTC()
	fmt.Println(newSunSet.Format("15:04:05"))

	DB.Create(&Currents)

	// if len(Currents.WeathersCurrent) > 0 {
	// 	DB. Current.WeathersCurrent[0]

	// 	// fmt.Println("First user:", firstUser)
	// } else {
	// 	fmt.Println("The Users slice is empty.")
	// }
}
