package database

import (
	"github.com/adeemgoogle/gowork/src/database/drivers"
	"github.com/adeemgoogle/gowork/src/database/drivers/postgres"
)

func New(config drivers.Config) (drivers.DataStore, error) {
	switch config.DSName {
	case "postgres":
		return postgres.New(config), nil
	default:
		return nil, drivers.ErrConfigNotSet
	}
}
