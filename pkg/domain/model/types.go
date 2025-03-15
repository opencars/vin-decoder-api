package model

import (
	"regexp"
)

var IsVIN = regexp.MustCompile(`^[A-HJ-NPR-Z0-9]{17}$`)

// Result is a union of information about vin-code and decoded vehicle.
type Result struct {
	VIN     *VIN            `json:"vin,omitempty"`
	Vehicle *Vehicle        `json:"vehicle,omitempty"`
	Error   *ProcesingError `json:"error,,omitempty"`
}

type ProcesingError struct {
	Messages []string `json:"messages"`
}

// Result is a union of information about vin-code and decoded vehicle.
type BulkResult struct {
	Results []Result `json:"results"`
}

// VIN represents detailed information about the VIN code.
type VIN struct {
	WMI string `json:"wmi"` // World manufacturer identifier.
	VDS string `json:"vds"` // Vehicle descriptor section.
	VIS string `json:"vis"` // Vehicle identifier section.
}

// Vehicle represent information about the decoded vehicle.
type Vehicle struct {
	Manufacturer string `json:"manufacturer"`
	Country      string `json:"country"`
	CountryUA    string `json:"country_ua"`
	Year         *uint  `json:"year,omitempty"`
	Region       Region `json:"region"`
	Check        bool   `json:"check"`
}

type Region string

const (
	Africa       Region = "Africa"
	Asia         Region = "Asia"
	Europe       Region = "Europe"
	NorthAmerica Region = "North America"
	Oceania      Region = "Oceania"
	SouthAmerica Region = "South America"
)
