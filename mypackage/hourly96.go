package mypackage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Hourlys struct {
	Hourlys    []Hourly   `json:"list"`
	CityHourly CityHourly `json:"city" gorm:"foreignKey:CityHourlyID"`
}
type CityHourly struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	TimeZone int64  `json:"timezone"`
}
type Hourly struct {
	WeathersHourly       []WeathersHourly     `json:"weather"`
	MainParametersHourly MainParametersHourly `json:"main"`
	Time                 int64                `json:"dt"`
	RainHourly           RainHourly           `json:"rain"`
}

type RainHourly struct {
	Rain float64 `json:"1h"`
}
type WeathersHourly struct {
	Weather string `json:"description"`
}
type MainParametersHourly struct {
	Temperature float64 `json:"temp"`
	Feels       float64 `json:"feels_like"`
}

func hourlyData(cityName string) {
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

	for i := 0; i < len(hourlys.Hourlys); i++ {
		fmt.Println("ID: " + strconv.Itoa(hourlys.CityHourly.Id))
		fmt.Println("City Name: " + hourlys.CityHourly.Name)
		fmt.Println("Weather Condition: " + hourlys.Hourlys[i].WeathersHourly[0].Weather)
		fmt.Println("Current Temperature: " + strconv.FormatFloat(hourlys.Hourlys[i].MainParametersHourly.Temperature, 'f', -2, 64))
		fmt.Println("Feels Like: " + strconv.FormatFloat(hourlys.Hourlys[i].MainParametersHourly.Feels, 'f', 0, 64))
		newSunRise := time.Unix(hourlys.Hourlys[i].Time+hourlys.CityHourly.TimeZone, 0).UTC()
		fmt.Println(newSunRise.Format("2006-01-02 15:04:05"))
		fmt.Println("Rain: " + strconv.FormatFloat(hourlys.Hourlys[i].RainHourly.Rain, 'f', -2, 64))

		fmt.Println("------------------------------------------")
	}
}
