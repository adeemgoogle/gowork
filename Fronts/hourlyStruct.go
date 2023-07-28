package Fronts

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
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

func GetHourly() {
	// Подключение к базе данных PostgreSQL (замените параметры подключения на ваши).
	dsn := "host=localhost user=postgres password=123 dbname=weather port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Ошибка при подключении к базе данных:", err)
		return
	}

	// Получение данных из базы данных с объединением таблиц и предварительной загрузкой связанных данных.
	var hourlys []Hourlys
	if err := db.Preload("Hourlys.WeathersHourly").Preload("Hourlys.MainParametersHourly").Find(&hourlys).Error; err != nil {
		fmt.Println("Ошибка при получении данных из базы данных:", err)
		return
	}

	// Преобразование данных в JSON формат.
	jsonData, err := json.MarshalIndent(hourlys, "", "    ")
	if err != nil {
		fmt.Println("Ошибка при преобразовании данных в JSON:", err)
		return
	}

	// Запись данных в файл.
	err = ioutil.WriteFile("hourly_info.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	fmt.Println("Данные успешно записаны в файл hourly_info.json")
}


