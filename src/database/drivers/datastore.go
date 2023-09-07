package drivers

import (
	"context"
	"github.com/adeemgoogle/gowork/src/model"
)

type Base interface {
	Connect(ctx context.Context) error
	Ping() error
	Close() error
	RunMigrations() error
	WeatherRepository() WeatherRepository
}

type DataStore interface {
	Base
}

type WeatherRepository interface {
	GetWeatherType(extId int) (model.WeatherType, error)
	CreateWeatherType(weatherType model.WeatherType) (model.WeatherType, error)

	GetCurrentByLocation(location string) (model.Current, error)
	SaveCurrent(current model.Current) (model.Current, error)

	GetHourliesByLocation(location string) ([]model.Hourly, error)
	CreateHourly(hourly model.Hourly) (model.Hourly, error)
	DeleteHourly(hourly model.Hourly) error

	GetClimatesByLocation(location string) ([]model.Climate, error)
	CreateClimate(hourly model.Climate) (model.Climate, error)
	DeleteClimate(hourly model.Climate) error
}
