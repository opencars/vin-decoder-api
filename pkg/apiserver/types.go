package apiserver

import (
	"github.com/opencars/vin-decoder-api/pkg/govin"
)

func NewResult(vin *govin.VIN) *Result {
	return &Result{
		VIN: VIN{
			WMI: vin.WMI(),
			VDS: vin.VDS(),
			VIS: vin.VIS(),
		},
		Vehicle: Vehicle{
			Manufacturer: vin.Manufacturer(),
			Country:      vin.Country(),
			Year:         vin.Year(),
			Region:       vin.Region(),
			Make:         vin.Make(),
			CheckDigit:   vin.Check(),
		},
	}
}

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
	Year         int          `json:"year"`
	Region       govin.Region `json:"region"`
	Make         string       `json:"make"`
	CheckDigit   bool         `json:"check_digit"`
}
