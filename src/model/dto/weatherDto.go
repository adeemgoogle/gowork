package dto

import (
	"time"
)

type WeatherDto struct {
	Current  CurrentDto   `json:"current"`
	Hourlies []HourlyDto  `json:"hourlies"`
	Climates []ClimateDto `json:"climates"`
}

type CurrentDto struct {
	Id           int64            `json:"id"`
	Location     string           `json:"location"`
	Temp         float64          `json:"temp"`
	FeelsLike    float64          `json:"feelsLike"`
	Date         time.Time        `json:"date"`
	Humidity     int              `json:"humidity"`
	Visibility   int              `json:"visibility"`
	WindSpeed    float64          `json:"windSpeed"`
	WindDeg      float64          `json:"windDeg"`
	Cloud        int              `json:"cloud"`
	Rain1h       float64          `json:"rain1H"`
	Rain3h       float64          `json:"rain3H"`
	Snow1h       float64          `json:"snow1H"`
	Snow3h       float64          `json:"snow3H"`
	WeatherTypes []WeatherTypeDto `json:"weatherTypes"`
}

type HourlyDto struct {
	Id           int64            `json:"id"`
	Location     string           `json:"location"`
	Temp         float64          `json:"temp"`
	FeelsLike    float64          `json:"feelsLike"`
	Date         time.Time        `json:"date"`
	WeatherTypes []WeatherTypeDto `json:"weatherTypes"`
}

type ClimateDto struct {
	Id           int64            `json:"id"`
	Location     string           `json:"location"`
	Temp         float64          `json:"temp"`
	FeelsLike    float64          `json:"feelsLike"`
	Date         time.Time        `json:"date"`
	WeatherTypes []WeatherTypeDto `json:"weatherTypes"`
}

type WeatherTypeDto struct {
	Id          int64  `json:"id"`
	ExtId       int    `json:"extId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
