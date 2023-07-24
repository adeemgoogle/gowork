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

type Dailys struct {
	ID uint        `gorm:"primaryKey;autoIncrement"`
	Daily []Daily `json:"list" gorm:"foreignKey: DailyID"`
	City   City    `json:"city" gorm:"foreignKey: CityDailyID"`
}
type City struct {
	CityDailyID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name"`
	TimeZone int64  `json:"timezone"`
}
type Daily struct {
	DailyID             uint                `gorm:"primaryKey;autoIncrement"`
	ParentDailyID       uint
	WeathersDaily       []WeathersDaily     `json:"weather" gorm:"foreignKey:WeathersDailyID"`
	MainParametersDaily MainParametersDaily `json:"temp"  gorm:"foreignKey:MainParametersDailyID"`
	Time                int64               `json:"dt"`
}
type WeathersDaily struct {
	WeathersDailyID uint `gorm:"primaryKey;autoIncrement"`
	Weather string `json:"description"`
}
type MainParametersDaily struct {
	MainParametersDailyID  uint  `gorm:"primaryKey;autoIncrement"`
	TemperatureMax float64 `json:"max"`
	TemperatureMin float64 `json:"min"`
}

func DailyData(cityName string, DB *gorm.DB) {
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

	var dailys Dailys

	json.Unmarshal(byteResult, &dailys)

	// for i := 0; i < len(dailys.Dailys); i++ {
	// 	// fmt.Println("ID: " + strconv.Itoa(dailys.City.Id))
	// 	fmt.Println("City Name: " + dailys.City.Name)
	// 	fmt.Println("Weather Condition: " + dailys.Dailys[i].WeathersDaily[0].Weather)
	// 	fmt.Println("Maximum Temperature: " + strconv.FormatFloat(dailys.Dailys[i].MainParametersDaily.TemperatureMax, 'f', -2, 64))
	// 	fmt.Println("Minimum Temperature: " + strconv.FormatFloat(dailys.Dailys[i].MainParametersDaily.TemperatureMin, 'f', -2, 64))
	// 	newSunRise := time.Unix(dailys.Dailys[i].Time+dailys.City.TimeZone, 0).UTC()
	// 	fmt.Println(newSunRise.Format("2006-01-02 15:04:05"))
	// 	fmt.Println("------------------------------------------")
	// }

	// DB.Create(&dailys)

}