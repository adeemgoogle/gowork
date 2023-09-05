package weather

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	httpClient "github.com/adeemgoogle/gowork/src/common/http"
	"github.com/adeemgoogle/gowork/src/config"
	"github.com/adeemgoogle/gowork/src/database/drivers"
	"github.com/adeemgoogle/gowork/src/model/response"
	"log"
)

type IService interface {
	GetCurrentData(ctx context.Context, config config.Config, location string) error
}

type Service struct {
	weatherRepo   drivers.WeatherRepository
	weatherClient *httpClient.Client
}

func NewService(weatherRepo drivers.WeatherRepository, weatherClient *httpClient.Client) *Service {
	return &Service{
		weatherRepo:   weatherRepo,
		weatherClient: weatherClient,
	}
}

func (s Service) GetCurrentData(ctx context.Context, config config.Config, location string) error {
	url := "/weather?q=" + location + "&appid=" + config.WeatherAppId
	resp, err := s.weatherClient.Get(ctx, url)
	if err != nil {
		return err
	}

	// преобразование ответа в JSON для отображения в консоли
	log.Println("Current response:", string(resp.Body))

	var response response.CurrentResponse
	err = json.Unmarshal(resp.Body, &response)
	if err != nil {
		return errors.New(fmt.Sprintf("Ошибка при разборе JSON: %s", err.Error()))
	}

	return nil
}
