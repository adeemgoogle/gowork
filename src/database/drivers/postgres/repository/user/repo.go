package user

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

func (r Repository) CreateUser(entity model.User) (model.User, error) {
	if err := r.db.Create(&entity).Error; err != nil {
		return model.User{}, err
	}

	return entity, nil
}

func (r Repository) GetUserByDeviceId(deviceId string) (model.User, error) {
	var entity model.User
	if err := r.db.Preload("Locations").Where("device_id = ?", deviceId).Find(&entity).Error; err != nil {
		return model.User{}, err
	}

	return entity, nil
}
