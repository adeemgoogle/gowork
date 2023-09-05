package model

type CurrentEntity struct {
	Id          int64  `gorm:"primaryKey, autoIncrement"`
	Type        string `gorm:"not null"`
	Description string `gorm:"not null"`
}
