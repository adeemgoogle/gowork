package modell

type Hourlys struct {
	ID         uint       `gorm:"primaryKey;autoIncrement"`
	Hourly     []Hourly   `json:"list"  gorm:"foreignKey: ParentID"`
	CityHourly CityHourly `json:"city" gorm:"foreignKey: CityHourlyID"`
}
type CityHourly struct {
	CityHourlyID int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string `json:"name"`
	TimeZone     int64  `json:"timezone"`
}
type Hourly struct {
	HourlyId             uint `gorm:"primaryKey;autoIncrement"`
	ParentID             uint
	WeathersHourly       []WeathersHourly     `json:"weather" gorm:"foreignKey:WeatherHourlyID"`
	MainParametersHourly MainParametersHourly `json:"main" gorm:"foreignKey:MainHourlyID"`
	Time                 int64                `json:"dt"`
	// RainHourly           RainHourly           `json:"rain"`
}

type RainHourly struct {
	Rain float64 `json:"1h"`
}
type WeathersHourly struct {
	WeatherHourlyID uint   `gorm:"primaryKey;autoIncrement"`
	Weather         string `json:"description"`
}
type MainParametersHourly struct {
	MainHourlyID uint    `gorm:"primaryKey;autoIncrement"`
	Temperature  float64 `json:"temp"`
	Feels        float64 `json:"feels_like"`
}
