package model

import "github.com/adeemgoogle/gowork/src/model/enum"

type User struct {
	Id        int64       `gorm:"primaryKey, autoIncrement"`
	DeviceId  string      `gorm:"not null; unique"`
	Gender    enum.Gender `gorm:"not null"`
	Locations []*Location `gorm:"many2many:user_locations;"`
}

type UserLocations struct {
	Id         int64 `gorm:"primaryKey, autoIncrement"`
	UserId     int64 `gorm:"not null"`
	LocationId int64 `gorm:"not null"`
}
