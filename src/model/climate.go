package model

import "time"

type Climate struct {
	Id           int64          `gorm:"primaryKey, autoIncrement"`                                             // идентификатор
	Location     string         `gorm:"not null; index:location_date_index"`                                   // локация
	Temp         float64        `gorm:"not null"`                                                              // температура
	FeelsLike    float64        `gorm:"not null"`                                                              // температура восприятие человеком погоды
	Date         time.Time      `gorm:"not null; index:location_date_index; type:timestamp without time zone"` // дата
	Timezone     string         `gorm:"not null"`                                                              // таймзона
	WeatherTypes []*WeatherType `gorm:"many2many:climate_weather_types"`                                       // список типа погод
}

type ClimateWeatherType struct {
	Id            int64 `gorm:"primaryKey, autoIncrement"`
	ClimateId     int64 `gorm:"not null"`
	WeatherTypeId int64 `gorm:"not null"`
}
