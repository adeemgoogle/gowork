package model

import "time"

type Climate struct {
	Id           int64          `gorm:"primaryKey, autoIncrement"`
	Location     string         `gorm:"not null"`
	Temp         float64        `gorm:"not null"`
	FeelsLike    float64        `gorm:"not null"`
	Date         time.Time      `gorm:"type:timestamp without time zone"`
	Timezone     string         `gorm:"not null"`
	WeatherTypes []*WeatherType `gorm:"many2many:climate_weather_types"`
}

type ClimateWeatherType struct {
	Id            int64 `gorm:"primaryKey, autoIncrement"`
	ClimateId     int64 `gorm:"not null"`
	WeatherTypeId int64 `gorm:"not null"`
}
