package weather

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/adeemgoogle/gowork/src/config"
	"github.com/adeemgoogle/gowork/src/model"
	"github.com/adeemgoogle/gowork/src/model/integ"
)

// checkAndGetHourlies - проверка и получения почасовых прогноз на 4 дня
func (s Service) checkAndGetHourlies(ctx context.Context, config config.Config, location string) ([]model.Hourly, error) {
	hourlies, err := s.weatherRepo.GetHourliesByLocation(location)
	if err != nil {
		return nil, err
	}

	// удаляет старые данные
	currentDate := time.Now()
	for _, hourly := range hourlies {
		if reconverTimezone(hourly.Date, hourly.Timezone).Before(currentDate) {
			err = s.weatherRepo.DeleteHourly(hourly)
			if err != nil {
				return nil, err
			}
		}
	}

	return s.updateHourliesData(ctx, config, location, hourlies)
}

// GetHourliesData - обновления почасовых прогноз на 4 дня
func (s Service) updateHourliesData(ctx context.Context, config config.Config, location string, hourlies []model.Hourly) ([]model.Hourly, error) {
	rsHourly, err := s.sendHourlyRequest(ctx, config, location)
	if err != nil {
		return nil, err
	}

	lastHourly := findHourlyWithMaxDate(hourlies)
	for _, r := range rsHourly.List {
		date := convertDate(r.Dt, rsHourly.City.Timezone)

		if lastHourly != nil {
			if date.After(lastHourly.Date) {
				weatherTypes, err := s.getWeatherTypes(ctx, config, r.Weather)
				if err != nil {
					return nil, err
				}

				entity := model.Hourly{
					Location:     location,
					Temp:         r.Main.Temp,
					FeelsLike:    r.Main.FeelsLike,
					Date:         date,
					WeatherTypes: weatherTypes,
				}
				_, err = s.weatherRepo.CreateHourly(entity)
				if err != nil {
					return nil, err
				}
			}
		} else {
			weatherTypes, err := s.getWeatherTypes(ctx, config, r.Weather)
			if err != nil {
				return nil, err
			}

			entity := model.Hourly{
				Location:     location,
				Temp:         r.Main.Temp,
				FeelsLike:    r.Main.FeelsLike,
				Date:         date,
				WeatherTypes: weatherTypes,
			}
			_, err = s.weatherRepo.CreateHourly(entity)
			if err != nil {
				return nil, err
			}
		}
	}

	entities, err := s.weatherRepo.GetHourliesByLocation(location)
	if err != nil {
		return nil, err
	}

	return entities, nil
}

// sendCurrentRequest - запрос почасовых прогноз на 4 дня
func (s Service) sendHourlyRequest(ctx context.Context, config config.Config, location string) (*integ.RsHourly, error) {
	url := "/forecast/hourly?q=" + location + "&units=metric&appid=" + config.WeatherAppId
	resp, err := s.weatherClient.Get(ctx, url)
	if err != nil {
		return nil, err
	}

	var rsHourly integ.RsHourly
	err = json.Unmarshal(resp.Body, &rsHourly)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Ошибка при разборе JSON: %s", err.Error()))
	}

	return &rsHourly, nil
}

// findHourlyWithMaxDate - получение последнего объекта по дате
func findHourlyWithMaxDate(hourlies []model.Hourly) *model.Hourly {
	if len(hourlies) == 0 {
		return nil
	}

	maxDateObject := &hourlies[0]
	maxDate := hourlies[0].Date
	for i := 1; i < len(hourlies); i++ {
		if hourlies[i].Date.After(maxDate) {
			maxDate = hourlies[i].Date
			maxDateObject = &hourlies[i]
		}
	}
	return maxDateObject
}
