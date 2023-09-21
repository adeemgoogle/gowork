package handler

import (
	"github.com/adeemgoogle/gowork/src/config"
	"github.com/adeemgoogle/gowork/src/service/location"
	"github.com/adeemgoogle/gowork/src/service/user"
	"github.com/adeemgoogle/gowork/src/service/weather"
)

type Handler struct {
	weatherService  weather.IService
	userService     user.IService
	locationService location.IService
	config          config.Config
}

func NewHandler(weatherService weather.IService, userService user.IService, locationService location.IService, config config.Config) *Handler {
	return &Handler{
		weatherService:  weatherService,
		userService:     userService,
		locationService: locationService,
		config:          config,
	}
}
