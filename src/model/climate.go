package model

import "time"

type Climate struct {
	Id           int64          `gorm:"primaryKey, autoIncrement"`
	Location     string         `gorm:"not null, index:location_date_index "`
	Temp         float64        `gorm:"not null"`
	FeelsLike    float64        `gorm:"not null"`
	Date         time.Time      `gorm:"not null, index:location_date_index"`
	WeatherTypes []*WeatherType `gorm:"many2many:climate_weather_types"`
	Sunrise 	int 			`gorm:"not null"`
}

type ClimateWeatherType struct {
	Id            int64 `gorm:"primaryKey, autoIncrement"`
	ClimateId     int64 `gorm:"not null"`
	WeatherTypeId int64 `gorm:"not null"`
}
