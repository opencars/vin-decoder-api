package sqlstore

import (
	"database/sql"
	"errors"

	"github.com/opencars/vin-decoder-api/pkg/domain/model"
)

type ManufacturerRepository struct {
	store *Store
}

func (r *ManufacturerRepository) Create(manufacturer *model.Manufacturer) error {
	_, err := r.store.db.NamedExec(
		`INSERT INTO manufacturers (
			wmi, name
		) VALUES (
			:wmi, :name
		)`,
		manufacturer,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *ManufacturerRepository) FindByWMI(wmi string) (*model.Manufacturer, error) {
	var manufacturer model.Manufacturer

	err := r.store.db.Get(&manufacturer,
		`SELECT wmi, name FROM manufacturers WHERE wmi = $1`,
		wmi,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, model.ErrManufacturerNotFound
	}

	if err != nil {
		return nil, err
	}

	return &manufacturer, nil
}
