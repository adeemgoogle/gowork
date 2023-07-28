package Fronts

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
)

type Dailys struct {
	ID uint        `gorm:"primaryKey;autoIncrement"`
	Daily []Daily `json:"list" gorm:"foreignKey: ParentDailyID"`
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


func GetDaily() {
	// Подключение к базе данных PostgreSQL (замените параметры подключения на ваши).
	dsn := "host=localhost user=postgres password=123 dbname=weather port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return
	}

	// Получение данных из базы данных с объединением таблиц и предварительной загрузкой связанных данных.
	var dailys []Dailys
	if err := db.Preload("Daily.WeathersDaily").Preload("Daily.MainParametersDaily").Find(&dailys).Error; err != nil {
		fmt.Println("Ошибка при получении данных из базы данных:", err)
		return
	}

	// Преобразование данных в JSON формат.
	jsonData, err := json.MarshalIndent(dailys, "", "    ")
	if err != nil {
		fmt.Println("Ошибка при преобразовании данных в JSON:", err)
		return
	}

	// Запись данных в файл.
	err = ioutil.WriteFile("daily.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	fmt.Println("Данные успешно записаны в файл daily.json")
}

