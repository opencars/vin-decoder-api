package service

import (
	"context"

	"github.com/opencars/schema"
	"github.com/opencars/seedwork"

	"github.com/opencars/vin-decoder-api/pkg/domain"
	"github.com/opencars/vin-decoder-api/pkg/domain/command"
	"github.com/opencars/vin-decoder-api/pkg/domain/model"
)

type CustomerService struct {
	repo domain.ManufacturerRepository
	p    schema.Producer
}

func NewCustomerService(repo domain.ManufacturerRepository, p schema.Producer) *CustomerService {
	return &CustomerService{
		repo: repo,
		p:    p,
	}
}

func (s *CustomerService) DecodeVIN(ctx context.Context, c *command.DecodeVIN) (*model.Result, error) {
	if err := seedwork.ProcessCommand(c); err != nil {
		return nil, err
	}

	vin := Parse(c.VIN)
	country := vin.Country()

	result := model.Result{
		VIN: &model.VIN{
			WMI: vin.WMI(),
			VDS: vin.VDS(),
			VIS: vin.VIS(),
		},
		Vehicle: &model.Vehicle{
			Manufacturer: vin.Manufacturer(s.repo),
			Year:         vin.Year(),
			Region:       vin.Region(),
			Check:        vin.Check(),
		},
	}

	if result.Vehicle != nil {
		result.Vehicle.Country = country.Name
		result.Vehicle.CountryUA = country.NameUA
	}

	if err := s.p.Produce(ctx, c.Event()); err != nil {
		return nil, err
	}

	return &result, nil
}
