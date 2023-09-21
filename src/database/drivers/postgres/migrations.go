package postgres

import (
	"github.com/adeemgoogle/gowork/src/database/drivers/postgres/migration"
	"github.com/adeemgoogle/gowork/src/model"
	"log"
)

func (p *Postgres) RunMigrations() error {
	err := p.db.AutoMigrate(model.WeatherType{}, model.Current{}, model.Hourly{}, model.Climate{},
		model.CurrentWeatherType{}, model.HourlyWeatherType{}, model.ClimateWeatherType{}, model.User{}, model.Location{})
	if err != nil {
		log.Fatal("failed to migrate db", err)
		return err
	}
	migration.InitDb(p.db)
	return nil
}
