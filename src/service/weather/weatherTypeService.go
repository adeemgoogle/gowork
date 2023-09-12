package weather

import (
	"context"
	"encoding/base64"
	httpClient "github.com/adeemgoogle/gowork/src/common/http"
	"github.com/adeemgoogle/gowork/src/model"
	"github.com/adeemgoogle/gowork/src/model/integ"
)

// getWeatherTypes - метод для получения типа погоды(достает по extId если есть, иначе создает новый)
func (s Service) getWeatherTypes(weathers []integ.RsWeather) ([]*model.WeatherType, error) {
	var list []*model.WeatherType
	for _, weather := range weathers {
		entity, err := s.weatherRepo.GetWeatherType(weather.Id)
		if err != nil {
			return nil, err
		}

		if entity.Id != 0 {
			list = append(list, &entity)
		} else {
			iconData, err := s.sendImageRequest(weather.Icon)
			if err != nil {
				return nil, err
			}

			entity, err = s.weatherRepo.CreateWeatherType(model.WeatherType{
				ExtId:       weather.Id,
				Name:        weather.Main,
				Description: weather.Description,
				Icon:        base64.StdEncoding.EncodeToString(iconData),
			})
			if err != nil {
				return nil, err
			}

			if entity.Id != 0 {
				list = append(list, &entity)
			}
		}
	}
	return list, nil
}

// sendClimateRequest - запрос прогноза климата на 30 дней
func (s Service) sendImageRequest(icon string) ([]byte, error) {
	weatherImageCli := httpClient.NewClient(s.config.WeatherImageURL)
	url := "/img/wn/" + icon + "@2x.png"
	resp, err := weatherImageCli.Get(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
