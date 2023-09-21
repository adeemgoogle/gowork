package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/adeemgoogle/gowork/src/config"
	"github.com/adeemgoogle/gowork/src/database/drivers"
	"github.com/adeemgoogle/gowork/src/model"
	"github.com/adeemgoogle/gowork/src/model/dto"
	"github.com/adeemgoogle/gowork/src/model/enum"
	"github.com/adeemgoogle/gowork/src/model/request"
)

type IService interface {
	CreateUser(ctx context.Context, deviceId string, rqUser request.RqUser) (*dto.UserDto, error)
	GetUser(ctx context.Context, deviceId string) (*dto.UserDto, error)
}

type Service struct {
	config       config.Config
	userRepo     drivers.UserRepository
	locationRepo drivers.LocationRepository
}

func NewService(config config.Config, userRepo drivers.UserRepository, locationRepo drivers.LocationRepository) *Service {
	return &Service{
		config:       config,
		userRepo:     userRepo,
		locationRepo: locationRepo,
	}
}

// CreateUser - сервис для создания профиля пользователя
func (s Service) CreateUser(ctx context.Context, deviceId string, rqUser request.RqUser) (*dto.UserDto, error) {
	gender, err := checkGender(rqUser.Gender)
	if err != nil {
		return nil, err
	}

	var locations []*model.Location
	for _, locationId := range rqUser.LocationIds {
		location, err := s.locationRepo.GetLocationById(locationId)
		if err != nil {
			return nil, err
		}

		locations = append(locations, &location)
	}

	user := model.User{
		DeviceId:  deviceId,
		Gender:    gender,
		Locations: locations,
	}
	entity, err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return buildUserDto(entity), nil
}

// GetUser - сервис для получения профиля пользователя
func (s Service) GetUser(ctx context.Context, deviceId string) (*dto.UserDto, error) {
	entity, err := s.userRepo.GetUserByDeviceId(deviceId)
	if err != nil {
		return nil, err
	}

	if entity.Id == 0 {
		return nil, errors.New(fmt.Sprintf("user not found by deviceId %s", deviceId))
	}

	return buildUserDto(entity), nil
}

func buildUserDto(entity model.User) *dto.UserDto {
	var locationsDto []*dto.LocationDto
	for _, e := range entity.Locations {
		locationDto := dto.LocationDto{
			Id:   e.Id,
			Name: e.Name,
		}
		locationsDto = append(locationsDto, &locationDto)
	}

	userDto := dto.UserDto{
		Id:        entity.Id,
		DeviceId:  entity.DeviceId,
		Gender:    entity.Gender,
		Locations: locationsDto,
	}
	return &userDto
}

func checkGender(gender enum.Gender) (enum.Gender, error) {
	switch gender {
	case enum.MALE, enum.FEMALE:
		return gender, nil
	default:
		return "", errors.New("incorrect gender")
	}
}
