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
