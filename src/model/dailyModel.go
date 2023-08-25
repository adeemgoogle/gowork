package modell

type Dailys struct {
	ID    uint    `gorm:"primaryKey;autoIncrement"`
	Daily []Daily `json:"list" gorm:"foreignKey: ParentDailyID"`
	City  City    `json:"city" gorm:"foreignKey: CityDailyID"`
}
type City struct {
	CityDailyID int    `gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name"`
	TimeZone    int64  `json:"timezone"`
}
type Daily struct {
	DailyID             uint `gorm:"primaryKey;autoIncrement"`
	ParentDailyID       uint
	WeathersDaily       []WeathersDaily     `json:"weather" gorm:"foreignKey:WeathersDailyID"`
	MainParametersDaily MainParametersDaily `json:"temp"  gorm:"foreignKey:MainParametersDailyID"`
	Time                int64               `json:"dt"`
}
type WeathersDaily struct {
	WeathersDailyID uint   `gorm:"primaryKey;autoIncrement"`
	Weather         string `json:"description"`
}
type MainParametersDaily struct {
	MainParametersDailyID uint    `gorm:"primaryKey;autoIncrement"`
	TemperatureMax        float64 `json:"max"`
	TemperatureMin        float64 `json:"min"`
}
