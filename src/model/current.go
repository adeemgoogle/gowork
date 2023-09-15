package model

import "time"

type Current struct {
	Id           int64          `gorm:"primaryKey, autoIncrement"`                  // идентификатор
	Location     string         `gorm:"not null; unique"`                           // локация
	Temp         float64        `gorm:"not null"`                                   // температура
	FeelsLike    float64        `gorm:"not null"`                                   // температура восприятие человеком погоды
	Date         time.Time      `gorm:"not null; type:timestamp without time zone"` // дата
	Timezone     string         `gorm:"not null"`                                   // таймзона
	Humidity     int            // влажность, %
	Visibility   int            // видимость, м
	WindSpeed    float64        // скорость ветра, м/с
	WindDeg      float64        // направление ветра, градус
	Cloud        int            // облачность, %
	Rain1h       float64        // объем дождя за последний 1 час, мм
	Rain3h       float64        // объем дождя за последний 3 часа, мм
	Snow1h       float64        // объем снега за последний 1 час, мм
	Snow3h       float64        // объем снега за последний 3 часа, мм
	WeatherTypes []*WeatherType `gorm:"many2many:current_weather_types;"` // список типа погод
}

type CurrentWeatherType struct {
	Id            int64 `gorm:"primaryKey, autoIncrement"`
	CurrentId     int64 `gorm:"not null"`
	WeatherTypeId int64 `gorm:"not null"`
}
