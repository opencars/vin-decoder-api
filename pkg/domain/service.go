package domain

import (
	"context"

	"github.com/opencars/vin-decoder-api/pkg/domain/command"
	"github.com/opencars/vin-decoder-api/pkg/domain/model"
)

// ManufacturerRepository is responsible for interaction with manufacturers entities.
type ManufacturerRepository interface {
	Create(manufacturer *model.Manufacturer) error
	FindByWMI(wmi string) (*model.Manufacturer, error)
}

type CustomerService interface {
	DecodeVIN(context.Context, *command.DecodeVIN) (*model.Result, error)
}

type InternalService interface {
	Decode(context.Context, *command.DecodeVINInternal) (*model.BulkResult, error)
}
