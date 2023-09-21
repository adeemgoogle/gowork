package postgres

import (
	"context"
	"github.com/adeemgoogle/gowork/src/database/drivers"
	"github.com/adeemgoogle/gowork/src/database/drivers/postgres/repository/location"
	"github.com/adeemgoogle/gowork/src/database/drivers/postgres/repository/user"
	"github.com/adeemgoogle/gowork/src/database/drivers/postgres/repository/weather"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

const driver = "postgres"

type Postgres struct {
	connURL string
	dbName  string

	db *gorm.DB

	connTimeout  time.Duration
	weatherRepo  drivers.WeatherRepository
	userRepo     drivers.UserRepository
	locationRepo drivers.LocationRepository
}

func New(config drivers.Config) *Postgres {
	return &Postgres{
		connURL: config.URL,
		dbName:  config.DBName,
	}
}
func (p *Postgres) Connect(ctx context.Context) error {
	var err error
	p.db, err = gorm.Open(postgres.Open(p.connURL), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) Ping() error {
	db, err := p.db.DB()
	if err != nil {
		return err
	}

	return db.Ping()
}

func (p *Postgres) Close() error {
	db, err := p.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func (p *Postgres) WeatherRepository() drivers.WeatherRepository {
	if p.weatherRepo == nil {
		p.weatherRepo = weather.NewRepository(p.db)
	}

	return p.weatherRepo
}

func (p *Postgres) UserRepository() drivers.UserRepository {
	if p.userRepo == nil {
		p.userRepo = user.NewRepository(p.db)
	}

	return p.userRepo
}

func (p *Postgres) LocationRepository() drivers.LocationRepository {
	if p.locationRepo == nil {
		p.locationRepo = location.NewRepository(p.db)
	}

	return p.locationRepo
}
