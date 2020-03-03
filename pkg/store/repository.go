package store

import (
	"github.com/opencars/vin-decoder-api/pkg/model"
)

// ManufacturerRepository is responsible for interaction with manufacturers entities.
type ManufacturerRepository interface {
	Create(manufacturer *model.Manufacturer) error
	FindByWMI(wmi string) (*model.Manufacturer, error)
}
