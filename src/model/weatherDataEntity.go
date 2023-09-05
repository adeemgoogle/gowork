package model

type WeatherDataEntity struct {
	Id        int64   `gorm:"primaryKey, autoIncrement"`
	Temp      float64 `gorm:"not null"`
	FeelsLike float64 `gorm:"not null"`
	TempMin   float64 `gorm:"not null"`
	TempMax   float64 `gorm:"not null"`
	Pressure  float64 `gorm:"not null"`
	Humidity  float64 `gorm:"not null"`
	WindSpeed int     `gorm:"not null"`
	WindDeg   int     `gorm:"not null"`
	Clouds    int     `gorm:"not null"`
}
