package sqlstore

import (
	"fmt"

	"github.com/opencars/vin-decoder-api/pkg/domain"

	"github.com/jmoiron/sqlx"

	"github.com/opencars/vin-decoder-api/pkg/config"
)

type Store struct {
	db *sqlx.DB

	manufacturerRepository *ManufacturerRepository
}

func (s *Store) Manufacturer() domain.ManufacturerRepository {
	if s.manufacturerRepository == nil {
		s.manufacturerRepository = &ManufacturerRepository{
			store: s,
		}
	}

	return s.manufacturerRepository
}

func New(settings *config.Database) (*Store, error) {
	info := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		settings.Host,
		settings.Port,
		settings.User,
		settings.Name,
		settings.SSLMode,
		settings.Password,
	)

	db, err := sqlx.Connect("postgres", info)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}
