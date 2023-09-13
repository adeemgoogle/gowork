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

// checkAndGetCurrent - проверка и получения текущих погодных данных
func (s Service) checkAndGetCurrent(ctx context.Context, config config.Config, location string) (*model.Current, error) {
	current, err := s.weatherRepo.GetCurrentByLocation(location)
	if err != nil {
		return nil, err
	}
	if time.Now().Sub(reconverTimezone(current.Date, current.Timezone)) < time.Hour {
		return &current, nil
	}
	return s.updateCurrentData(ctx, config, location, current)
}

// GetCurrentData - обновления текущих погодных данных
func (s Service) updateCurrentData(ctx context.Context, config config.Config, location string, current model.Current) (*model.Current, error) {
	rsCurrent, err := s.sendCurrentRequest(ctx, config, location)
	if err != nil {
		return nil, err
	}

	weatherTypes, err := s.getWeatherTypes(ctx, config, rsCurrent.Weather)
	if err != nil {
		return nil, err
	}
	timezone := converTimezone(rsCurrent.Timezone)
	date := convertDate(rsCurrent.Dt, rsCurrent.Timezone)
	entity := model.Current{
		Id:           current.Id,
		Location:     location,
		Temp:         rsCurrent.Main.Temp,
		FeelsLike:    rsCurrent.Main.FeelsLike,
		Date:         date,
		Timezone:     timezone,
		WeatherTypes: weatherTypes,
	}
	entity, err = s.weatherRepo.SaveCurrent(entity)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

// sendCurrentRequest - запрос текущих погодных данных
func (s Service) sendCurrentRequest(ctx context.Context, config config.Config, location string) (*integ.RsCurrent, error) {
	url := "/weather?q=" + location + "&units=metric&appid=" + config.WeatherAppId
	resp, err := s.weatherClient.Get(ctx, url)
	if err != nil {
		return nil, err
	}

	var rsCurrent integ.RsCurrent
	err = json.Unmarshal(resp.Body, &rsCurrent)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Ошибка при разборе JSON: %s", err.Error()))
	}

	return &rsCurrent, nil
}
