package model

type Location struct {
	Id        int64  `gorm:"primaryKey, autoIncrement"`
	Name      string `gorm:"not null"`
	CityId    string
	CountryId string
	RegionId  string
}
