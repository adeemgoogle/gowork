package postgres

import (
	"github.com/adeemgoogle/gowork/src/model"
	"log"
)

func (p *Postgres) RunMigrations() error {
	err := p.db.AutoMigrate(model.WeatherType{}, model.Current{}, model.Hourly{}, model.Climate{},
		model.CurrentWeatherType{}, model.HourlyWeatherType{}, model.ClimateWeatherType{})
	if err != nil {
		log.Fatal("failed to migrate db", err)
		return err
	}
	return nil
}
