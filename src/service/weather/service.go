package weather

import (
	"context"
	httpClient "github.com/adeemgoogle/gowork/src/common/http"
	"github.com/adeemgoogle/gowork/src/config"
	"github.com/adeemgoogle/gowork/src/database/drivers"
	"github.com/adeemgoogle/gowork/src/model"
	"github.com/adeemgoogle/gowork/src/model/dto"
	"time"
)

type IService interface {
	GetAllWeatherData(ctx context.Context, config config.Config, location string) (*dto.WeatherDto, error)
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

// GetAllWeatherData - сервис для получения всех данных о погоде
func (s Service) GetAllWeatherData(ctx context.Context, config config.Config, location string) (*dto.WeatherDto, error) {
	current, err := s.checkAndGetCurrent(ctx, config, location)
	if current == nil || err != nil {
		return nil, err
	}

	hourlies, err := s.checkAndGetHourlies(ctx, config, location)
	if err != nil {
		return nil, err
	}

	climates, err := s.checkAndGetClimates(ctx, config, location)
	if err != nil {
		return nil, err
	}

	return buildWeatherDto(*current, hourlies, climates), nil
}

// buildWeatherDto - собирает данные о погоде для ответа
func buildWeatherDto(current model.Current, hourlies []model.Hourly, climates []model.Climate) *dto.WeatherDto {
	currentDto := dto.CurrentDto{
		Id:           current.Id,
		Location:     current.Location,
		Temp:         current.Temp,
		FeelsLike:    current.FeelsLike,
		Date:         current.Date,
		WindSpeed: 	 current.WindSpeed,
		WindDeg: current.WindDeg,
		WeatherTypes: buildWeatherTypesDto(current.WeatherTypes),
	}

	var hourliesDto []dto.HourlyDto
	for _, hourly := range hourlies {
		hourlyDto := dto.HourlyDto{
			Id:           hourly.Id,
			Location:     hourly.Location,
			Temp:         hourly.Temp,
			FeelsLike:    hourly.FeelsLike,
			Date:         hourly.Date,
			WeatherTypes: buildWeatherTypesDto(hourly.WeatherTypes),
		}

		hourliesDto = append(hourliesDto, hourlyDto)
	}

	var climatesDto []dto.ClimateDto
	for _, climate := range climates {
		climateDto := dto.ClimateDto{
			Id:           climate.Id,
			Location:     climate.Location,
			Temp:         climate.Temp,
			FeelsLike:    climate.FeelsLike,
			Date:         climate.Date,
			WeatherTypes: buildWeatherTypesDto(climate.WeatherTypes),
		}

		climatesDto = append(climatesDto, climateDto)
	}

	return &dto.WeatherDto{
		Current:  currentDto,
		Hourlies: hourliesDto,
		Climates: climatesDto,
	}
}

// buildWeatherTypesDto - собирает данные о типе погоды для ответа
func buildWeatherTypesDto(weatherTypes []*model.WeatherType) []dto.WeatherTypeDto {
	var weatherTypesDto []dto.WeatherTypeDto
	for _, weatherType := range weatherTypes {
		weatherTypeDto := dto.WeatherTypeDto{
			Id:          weatherType.Id,
			ExtId:       weatherType.ExtId,
			Name:        weatherType.Name,
			Description: weatherType.Description,
			Icon:        weatherType.Icon,
		}

		weatherTypesDto = append(weatherTypesDto, weatherTypeDto)
	}
	return weatherTypesDto
}

// convertDate - метод для преоброзования даты
func convertDate(dateUnix int64, timeZone int) time.Time {
	// временное местоположение, которое учитывает часовой пояс (в секундах от UTC)
	timeLocation := time.FixedZone("Custom Time Zone", timeZone)

	// преобразует дату Unix во временную точку с учетом временной зоны
	return time.Unix(dateUnix, 0).In(timeLocation)
}
