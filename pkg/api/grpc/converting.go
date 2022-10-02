package grpc

import (
	"github.com/opencars/grpc/pkg/common"
	"github.com/opencars/grpc/pkg/vin_decoding"

	"github.com/opencars/vin-decoder-api/pkg/domain/model"
)

func ResultItemFromDomain(r *model.Result) *vin_decoding.DecodeResultItem {
	item := vin_decoding.DecodeResultItem{}

	if r.Error != nil {
		item.Error = &common.Error{
			Messages: r.Error.Messages,
		}
	}

	if r.Vehicle != nil {
		vehicle := &vin_decoding.Vehicle{
			Check:        r.Vehicle.Check,
			Country:      r.Vehicle.Country,
			Manufacturer: r.Vehicle.Manufacturer,
			Region:       string(r.Vehicle.Region),
		}

		if r.Vehicle.Year != nil {
			vehicle.Year = uint32(*r.Vehicle.Year)
		}

		item.Vehicle = vehicle
	}

	if r.VIN != nil {
		vin := &vin_decoding.DecodedVIN{
			Vds: r.VIN.VDS,
			Vis: r.VIN.VIS,
			Wmi: r.VIN.WMI,
		}

		item.DecodedVin = vin
	}

	return &item
}
