package migration

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/adeemgoogle/gowork/src/model"
	"github.com/adeemgoogle/gowork/src/model/integ"
	"gorm.io/gorm"
	"os"
)

func InitDb(db *gorm.DB) error {
	if err := initLocations(db); err != nil {
		return err
	}
	return nil
}

func initLocations(db *gorm.DB) error {
	var count int64
	db.Model(&model.Location{}).Count(&count)
	if count != 0 {
		return nil
	}

	locations, err := getLocationsFromFile()
	if err != nil {
		return err
	}

	tx := db.Begin()
	for _, vs := range locations {
		if err := tx.Create(&vs).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func getLocationsFromFile() ([]model.Location, error) {
	file, err := os.Open("locations.json")
	if err != nil {
		return nil, errors.New(fmt.Sprintln("Ошибка при открытии файла:", err))
	}
	defer file.Close()

	var rsLocations integ.RsLocations
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&rsLocations); err != nil {
		return nil, errors.New(fmt.Sprintln("Ошибка при декодировании JSON:", err))
	}

	var locations []model.Location
	for _, r := range rsLocations.Location {
		location := model.Location{
			Name:      r.Name,
			CityId:    r.CityId,
			CountryId: r.CountryId,
			RegionId:  r.RegionId,
		}
		locations = append(locations, location)
	}

	return locations, nil
}
