package drivers

import (
	"context"
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
}
