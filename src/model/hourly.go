package model

import (
	"time"
)

type Hourly struct {
	Id           int64          `gorm:"primaryKey, autoIncrement"`
	Location     string         `gorm:"not null, index:location_date_index"`
	Temp         float64        `gorm:"not null"`
	FeelsLike    float64        `gorm:"not null"`
	Date         time.Time      `gorm:"not null, index:location_date_index"`
	WeatherTypes []*WeatherType `gorm:"many2many:hourly_weather_types"`
}

type HourlyWeatherType struct {
	Id            int64 `gorm:"primaryKey, autoIncrement"`
	HourlyId      int64 `gorm:"not null"`
	WeatherTypeId int64 `gorm:"not null"`
}
