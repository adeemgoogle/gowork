package location

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

func (r Repository) GetLocations() ([]model.Location, error) {
	var entities []model.Location
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}

	return entities, nil
}

func (r Repository) GetLocationById(id int64) (model.Location, error) {
	var entity model.Location
	if err := r.db.Where("id = ?", id).Find(&entity).Error; err != nil {
		return model.Location{}, err
	}

	return entity, nil
}
