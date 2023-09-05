package postgres

import (
	"log"
)

func (p *Postgres) RunMigrations() error {
	err := p.db.AutoMigrate()
	if err != nil {
		log.Fatal("failed to migrate db", err)
		return err
	}
	return nil
}
