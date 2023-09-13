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
	WeatherTypes []WeatherTypeDto `json:"weatherTypes"`
	WindSpeed 	 float64 	  	  `json:"windSpeed"`
	WindDeg 	 float64 		  `json:"windDeg"`
	Visibility   int 		      `json:"visibility"`
}

type HourlyDto struct {
	Id           int64            `json:"id"`
	Location     string           `json:"location"`
	Temp         float64          `json:"temp"`
	FeelsLike    float64          `json:"feelsLike"`
	Date         time.Time        `json:"date"`
	WeatherTypes []WeatherTypeDto `json:"weatherTypes"`
	Sunrise 	int 				`json:"sunrise"`
	Sunset 		int  				`json:"sunset"`
}

type ClimateDto struct {
	Id           int64            `json:"id"`
	Location     string           `json:"location"`
	Temp         float64          `json:"temp"`
	FeelsLike    float64          `json:"feelsLike"`
	Date         time.Time        `json:"date"`
	WeatherTypes []WeatherTypeDto `json:"weatherTypes"`
	Sunrise 	 int64 			  `json:"sunrise"`
	Sunset       int64 			  `json:"sunset"`
	Humidity     int 			  `json:"humidity"`
	Pressure	 int 			  `json:"pressure"`
	MinTemp		float64 	  	  `json:"MinTemp"`
	MaxTemp 	float64 	  	  `json:"MaxTemp"`
}

type WeatherTypeDto struct {
	Id          int64  `json:"id"`
	ExtId       int    `json:"extId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
