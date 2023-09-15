package model

import "time"

type Current struct {
	Id           int64          `gorm:"primaryKey, autoIncrement"`
	Location     string         `gorm:"not null; unique"`
	Temp         float64        `gorm:"not null"`
	FeelsLike    float64        `gorm:"not null"`
	Date         time.Time      `gorm:"type:timestamp without time zone"`
	Timezone     string         `gorm:"not null"`
	WeatherTypes []*WeatherType `gorm:"many2many:current_weather_types;"`
}

type CurrentWeatherType struct {
	Id            int64 `gorm:"primaryKey, autoIncrement"`
	CurrentId     int64 `gorm:"not null"`
	WeatherTypeId int64 `gorm:"not null"`
}
