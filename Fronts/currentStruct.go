package Fronts

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
)

type Current struct {
	Id             int                   `json:"id" gorm:"primaryKey"`
	NameCurrent           string                `json:"name"`
	WeathersCurrent       []WeathersCurrent     `json:"weather" gorm:"foreignKey: WeatherID"`
	MainParametersCurrent *MainParametersCurrent `json:"main" gorm:"foreignKey: CurrentID"`
	WindCurrnet           WindCurrent           `json:"wind" gorm:"foreignKey: WindID"`
	InfoSunCurrent        *InfoSunCurrent        `json:"sys" gorm:"foreignKey:SunsetID"`
	TimeZone              int64                 `json:"timezone"`
	CloudsCurrent         *CloudsCurrent         `json:"clouds" gorm:"foreignKey: CloudsId"`
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

func GetCurrent() {
	// Подключение к базе данных PostgreSQL (замените параметры подключения на ваши).
	dsn := "host=localhost user=postgres password=123 dbname=weather port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return
	}

	// Получение данных из базы данных с объединением таблиц и предварительной загрузкой связанных данных.
	var currents []Current
	if err := db.Preload("WeathersCurrent").Preload("MainParametersCurrent").Preload("WindCurrnet").Preload("InfoSunCurrent").Preload("CloudsCurrent").Find(&currents).Error; err != nil {
		fmt.Println("Ошибка при получении данных из базы данных:", err)
		return
	}

	// Преобразование данных в JSON формат.
	jsonData, err := json.MarshalIndent(currents, "", "    ")
	if err != nil {
		fmt.Println("Ошибка при преобразовании данных в JSON:", err)
		return
	}

	// Запись данных в файл.
	err = ioutil.WriteFile("sun_info.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	fmt.Println("Данные успешно записаны в файл sun_info.json")
}

