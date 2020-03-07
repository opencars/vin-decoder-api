package apiserver

import (
	"github.com/opencars/vin-decoder-api/pkg/govin"
	"github.com/opencars/vin-decoder-api/pkg/store"
)

// Result is a union of information about vin-code and decoded vehicle.
type Result struct {
	VIN     VIN     `json:"vin"`
	Vehicle Vehicle `json:"vehicle"`
}

// VIN represents detailed information about the VIN code.
type VIN struct {
	WMI string `json:"wmi"` // World manufacturer identifier.
	VDS string `json:"vds"` // Vehicle descriptor section.
	VIS string `json:"vis"` // Vehicle identifier section.
}

// Vehicle represent information about the decoded vehicle.
type Vehicle struct {
	Manufacturer string       `json:"manufacturer"`
	Country      string       `json:"country"`
	Year         *uint        `json:"year,omitempty"`
	Region       govin.Region `json:"region"`
	Check        bool         `json:"check"`
}

func NewResult(store store.Store, vin *govin.VIN) *Result {
	return &Result{
		VIN: VIN{
			WMI: vin.WMI(),
			VDS: vin.VDS(),
			VIS: vin.VIS(),
		},
		Vehicle: Vehicle{
			Manufacturer: vin.Manufacturer(store),
			Country:      vin.Country(),
			Year:         vin.Year(),
			Region:       vin.Region(),
			Check:        vin.Check(),
		},
	}
}
