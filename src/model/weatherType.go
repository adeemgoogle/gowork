package model

type WeatherType struct {
	Id          int64  `gorm:"primaryKey, autoIncrement"`
	ExtId       int    `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Icon        string `gorm:"not null"`
}
