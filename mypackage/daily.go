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

type Dailys struct {
	Dailys []Daily `json:"list"`
	City   City    `json:"city"`
}
type City struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	TimeZone int64  `json:"timezone"`
}
type Daily struct {
	WeathersDaily       []WeathersDaily     `json:"weather"`
	MainParametersDaily MainParametersDaily `json:"temp"`
	Time                int64               `json:"dt"`
}
type WeathersDaily struct {
	Weather string `json:"description"`
}
type MainParametersDaily struct {
	TemperatureMax float64 `json:"max"`
	TemperatureMin float64 `json:"min"`
}

func main() {
	resp, err := http.Get("https://pro.openweathermap.org/data/2.5/forecast/daily?q=Almaty&appid=51e51b22fb137270e2e89bd2bc7c4acc&units=metric")
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

	for i := 0; i < len(dailys.Dailys); i++ {
		fmt.Println("ID: " + strconv.Itoa(dailys.City.Id))
		fmt.Println("City Name: " + dailys.City.Name)
		fmt.Println("Weather Condition: " + dailys.Dailys[i].WeathersDaily[0].Weather)
		fmt.Println("Maximum Temperature: " + strconv.FormatFloat(dailys.Dailys[i].MainParametersDaily.TemperatureMax, 'f', -2, 64))
		fmt.Println("Minimum Temperature: " + strconv.FormatFloat(dailys.Dailys[i].MainParametersDaily.TemperatureMin, 'f', -2, 64))
		newSunRise := time.Unix(dailys.Dailys[i].Time+dailys.City.TimeZone, 0).UTC()
		fmt.Println(newSunRise.Format("2006-01-02 15:04:05"))
		fmt.Println("------------------------------------------")
	}
}