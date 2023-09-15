package weather

import (
	"context"
	"strconv"
	"time"

	httpClient "github.com/adeemgoogle/gowork/src/common/http"
	"github.com/adeemgoogle/gowork/src/config"
	"github.com/adeemgoogle/gowork/src/database/drivers"
	"github.com/adeemgoogle/gowork/src/model"
	"github.com/adeemgoogle/gowork/src/model/dto"
)

type IService interface {
	GetAllWeatherData(ctx context.Context, location string) (*dto.WeatherDto, error)
}

type Service struct {
	config        config.Config
	weatherRepo   drivers.WeatherRepository
	weatherClient *httpClient.Client
}

func NewService(config config.Config, weatherRepo drivers.WeatherRepository, weatherClient *httpClient.Client) *Service {
	return &Service{
		config:        config,
		weatherRepo:   weatherRepo,
		weatherClient: weatherClient,
	}
}

// GetAllWeatherData - сервис для получения всех данных о погоде
func (s Service) GetAllWeatherData(ctx context.Context, location string) (*dto.WeatherDto, error) {
	var current *model.Current
	var hourlies []model.Hourly
	var climates []model.Climate
	var currentErr, hourliesErr, climatesErr error

	// Запускаем goroutines для каждого метода
	go func() {
		current, currentErr = s.checkAndGetCurrent(ctx, location)
	}()

	go func() {
		hourlies, hourliesErr = s.checkAndGetHourlies(ctx, location)
	}()

	go func() {
		climates, climatesErr = s.checkAndGetClimates(ctx, location)
	}()

	// Ожидаем завершения всех goroutines
	for i := 0; i < 3; i++ {
		<-time.After(time.Second * 5) // Ограничиваем время ожидания каждой goroutine
	}

	// Проверяем ошибки после завершения всех goroutines
	if currentErr != nil {
		return nil, currentErr
	}

	if hourliesErr != nil {
		return nil, hourliesErr
	}

	if climatesErr != nil {
		return nil, climatesErr
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
		Timezone:     current.Timezone,
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

// convertTimezone - метод для преоброзования таймзоны
func convertTimezone(timeZone int) string {
	//преоброзуем timezone в string
	time := strconv.Itoa(timeZone / 3600)
	return "GMT " + time
}

// reconvertTimezone - метод преоброзования времени обновления в таймзону системы
func reconvertTimezone(dateUTC time.Time, timeZone string) time.Time {
	switch {
	case len(timeZone) == 5:
		//преобразуем timezone в int
		intValue, _ := strconv.Atoi(string(timeZone[4]))
		//пребразуем в часовой формат
		utcTimeZone := time.Hour * time.Duration(intValue) * -1
		//получаем отрезок времени с нужной временной зоной
		result := dateUTC.Add(utcTimeZone)
		return result
	case len(timeZone) == 6:
		//преобразуем timezone в int
		intValue, _ := strconv.Atoi(string(timeZone[5]))
		//пребразуем в часовой формат
		utcTimeZone := time.Hour * time.Duration(intValue)
		//получаем отрезок времени с нужной временной зоной
		result := dateUTC.Add(utcTimeZone)
		return result
	default:
		return dateUTC
	}
}
