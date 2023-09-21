package model

import "github.com/adeemgoogle/gowork/src/model/enum"

type User struct {
	Id       int64       `gorm:"primaryKey, autoIncrement"`
	DeviceId string      `gorm:"not null; unique"`
	Gender   enum.Gender `grom:"not null"`
	Location string      `grom:"not null"`
}
