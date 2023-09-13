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

// checkAndGetClimates - проверка и получения прогноза климата на 30 дней
func (s Service) checkAndGetClimates(ctx context.Context, config config.Config, location string) ([]model.Climate, error) {
	climates, err := s.weatherRepo.GetClimatesByLocation(location)
	if err != nil {
		return nil, err
	}

	// удаляет старые данные
	currentDate := time.Now()
	for _, climate := range climates {
		if reconverTimezone(climate.Date, climate.Timezone).Before(currentDate) {
			err = s.weatherRepo.DeleteClimate(climate)
			if err != nil {
				return nil, err
			}
		}
	}

	return s.updateClimatesData(ctx, config, location, climates)
}

// updateClimatesData - обновления прогноза климата на 30 дней
func (s Service) updateClimatesData(ctx context.Context, config config.Config, location string, climates []model.Climate) ([]model.Climate, error) {
	rsClimate, err := s.sendClimateRequest(ctx, config, location)
	if err != nil {
		return nil, err
	}

	lastClimate := findClimateWithMaxDate(climates)
	for _, r := range rsClimate.List {
		date := convertDate(r.Dt, rsClimate.City.Timezone)

		if lastClimate != nil {
			if date.After(lastClimate.Date) {
				weatherTypes, err := s.getWeatherTypes(ctx, config, r.Weather)
				if err != nil {
					return nil, err
				}

				entity := model.Climate{
					Location:     location,
					Temp:         r.Temp.Day,
					FeelsLike:    r.FeelsLike.Day,
					Date:         date,
					WeatherTypes: weatherTypes,
				}
				_, err = s.weatherRepo.CreateClimate(entity)
				if err != nil {
					return nil, err
				}
			}
		} else {
			weatherTypes, err := s.getWeatherTypes(ctx, config, r.Weather)
			if err != nil {
				return nil, err
			}

			entity := model.Climate{
				Location:     location,
				Temp:         r.Temp.Day,
				FeelsLike:    r.FeelsLike.Day,
				Date:         date,
				WeatherTypes: weatherTypes,
			}
			_, err = s.weatherRepo.CreateClimate(entity)
			if err != nil {
				return nil, err
			}
		}
	}

	entities, err := s.weatherRepo.GetClimatesByLocation(location)
	if err != nil {
		return nil, err
	}

	return entities, nil
}

// sendClimateRequest - запрос прогноза климата на 30 дней
func (s Service) sendClimateRequest(ctx context.Context, config config.Config, location string) (*integ.RsClimate, error) {
	url := "/forecast/climate?q=" + location + "&units=metric&appid=" + config.WeatherAppId
	resp, err := s.weatherClient.Get(ctx, url)
	if err != nil {
		return nil, err
	}

	var rsClimate integ.RsClimate
	err = json.Unmarshal(resp.Body, &rsClimate)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Ошибка при разборе JSON: %s", err.Error()))
	}

	return &rsClimate, nil
}

// findClimateWithMaxDate - получение последнего объекта по дате
func findClimateWithMaxDate(climates []model.Climate) *model.Climate {
	if len(climates) == 0 {
		return nil
	}

	maxDateObject := &climates[0]
	maxDate := climates[0].Date
	for i := 1; i < len(climates); i++ {
		if climates[i].Date.After(maxDate) {
			maxDate = climates[i].Date
			maxDateObject = &climates[i]
		}
	}
	return maxDateObject
}
