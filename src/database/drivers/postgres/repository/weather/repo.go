package weather

import (
	"github.com/adeemgoogle/gowork/src/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) GetWeatherType(extId int) (model.WeatherType, error) {
	var entity model.WeatherType
	if err := r.db.Where("ext_id = ?", extId).Find(&entity).Error; err != nil {
		return model.WeatherType{}, err
	}
	return entity, nil
}

func (r Repository) CreateWeatherType(entity model.WeatherType) (model.WeatherType, error) {
	if err := r.db.Create(&entity).Error; err != nil {
		return model.WeatherType{}, err
	}

	return entity, nil
}

func (r Repository) GetCurrentByLocation(location string) (model.Current, error) {
	var entity model.Current
	if err := r.db.Preload("WeatherTypes").Where("location = ?", location).Find(&entity).Error; err != nil {
		return model.Current{}, err
	}
	return entity, nil
}

func (r Repository) SaveCurrent(entity model.Current) (model.Current, error) {
	if err := r.db.Save(&entity).Error; err != nil {
		return model.Current{}, err
	}

	return entity, nil
}

func (r Repository) GetHourliesByLocation(location string) ([]model.Hourly, error) {
	var entities []model.Hourly
	if err := r.db.Preload("WeatherTypes").Where("location = ?", location).Order("date").Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r Repository) CreateHourly(entity model.Hourly) (model.Hourly, error) {
	if err := r.db.Create(&entity).Error; err != nil {
		return model.Hourly{}, err
	}

	return entity, nil
}

func (r Repository) DeleteHourly(entity model.Hourly) error {
	var hourlyWeatherType model.HourlyWeatherType
	if err := r.db.Where("hourly_id = ?", entity.Id).Delete(hourlyWeatherType).Error; err != nil {
		return err
	}

	if err := r.db.Unscoped().Delete(entity).Error; err != nil {
		return err
	}
	return nil
}

func (r Repository) GetClimatesByLocation(location string) ([]model.Climate, error) {
	var entities []model.Climate
	if err := r.db.Preload("WeatherTypes").Where("location = ?", location).Order("date").Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r Repository) CreateClimate(entity model.Climate) (model.Climate, error) {
	if err := r.db.Create(&entity).Error; err != nil {
		return model.Climate{}, err
	}

	return entity, nil
}

func (r Repository) DeleteClimate(entity model.Climate) error {
	var climateWeatherType model.ClimateWeatherType
	if err := r.db.Where("climate_id = ?", entity.Id).Delete(climateWeatherType).Error; err != nil {
		return err
	}

	if err := r.db.Unscoped().Delete(entity).Error; err != nil {
		return err
	}
	return nil
}
