package service

import (
	"context"
	"fmt"

	"github.com/opencars/seedwork"

	"github.com/opencars/vin-decoder-api/pkg/domain"
	"github.com/opencars/vin-decoder-api/pkg/domain/command"
	"github.com/opencars/vin-decoder-api/pkg/domain/model"
)

type InternalService struct {
	repo domain.ManufacturerRepository
}

func NewInternalService(repo domain.ManufacturerRepository) *InternalService {
	return &InternalService{
		repo: repo,
	}
}

func (s *InternalService) Decode(ctx context.Context, c *command.DecodeVINInternal) (*model.BulkResult, error) {
	results := make([]model.Result, 0, len(c.Items))

	for _, v := range c.Items {
		if err := v.Validate(); err != nil {
			msgs := make([]string, 0)
			for k, vv := range seedwork.ErrorMessages("item", err) {
				for _, v := range vv {
					msgs = append(msgs, fmt.Sprintf("%s.%s", k, v))
				}
			}

			results = append(results, model.Result{
				Error: &model.ProcesingError{
					Messages: msgs,
				},
			})

			continue
		}

		vin := Parse(v.VIN)
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

		results = append(results, result)
	}

	return &model.BulkResult{Results: results}, nil
}
