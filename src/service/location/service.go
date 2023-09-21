package location

import (
	"context"
	"github.com/adeemgoogle/gowork/src/config"
	"github.com/adeemgoogle/gowork/src/database/drivers"
	"github.com/adeemgoogle/gowork/src/model/dto"
)

type IService interface {
	GetLocations(ctx context.Context) ([]dto.LocationDto, error)
}

type Service struct {
	config       config.Config
	locationRepo drivers.LocationRepository
}

func NewService(config config.Config, locationRepo drivers.LocationRepository) *Service {
	return &Service{
		config:       config,
		locationRepo: locationRepo,
	}
}

// GetLocations - сервис для получения списка всех локации
func (s Service) GetLocations(ctx context.Context) ([]dto.LocationDto, error) {
	entities, err := s.locationRepo.GetLocations()
	if err != nil {
		return nil, err
	}

	var locationsDto []dto.LocationDto
	for _, entity := range entities {
		locationDto := dto.LocationDto{
			Id:   entity.Id,
			Name: entity.Name,
		}
		locationsDto = append(locationsDto, locationDto)
	}
	return locationsDto, nil
}
