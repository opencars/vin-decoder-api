package sqlstore

import (
	"fmt"

	"github.com/opencars/vin-decoder-api/pkg/store"

	"github.com/jmoiron/sqlx"

	"github.com/opencars/vin-decoder-api/pkg/config"
)

type Store struct {
	db *sqlx.DB

	manufacturerRepository *ManufacturerRepository
}

func (s *Store) Manufacturer() store.ManufacturerRepository {
	if s.manufacturerRepository == nil {
		s.manufacturerRepository = &ManufacturerRepository{
			store: s,
		}
	}

	return s.manufacturerRepository
}

func New(conf *config.Database) (*Store, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		conf.Host, conf.Port, conf.User, conf.Name, conf.SSLMode, conf.Password,
	)

	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}
