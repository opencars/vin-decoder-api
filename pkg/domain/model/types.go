package model

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
	Manufacturer string `json:"manufacturer"`
	Country      string `json:"country"`
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
