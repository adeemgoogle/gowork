package model

import (
	"time"
)

type Hourly struct {
	Id           int64          `gorm:"primaryKey, autoIncrement"`
	Location     string         `gorm:"not null"`
	Temp         float64        `gorm:"not null"`
	FeelsLike    float64        `gorm:"not null"`
	Date         time.Time      `gorm:"type:timestamp without time zone"`
	Timezone     string         `gorm:"not null"`
	WeatherTypes []*WeatherType `gorm:"many2many:hourly_weather_types"`
}

type HourlyWeatherType struct {
	Id            int64 `gorm:"primaryKey, autoIncrement"`
	HourlyId      int64 `gorm:"not null"`
	WeatherTypeId int64 `gorm:"not null"`
}
