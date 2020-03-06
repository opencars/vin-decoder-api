package model

// Manufacturer represents entities from manufacturer table.
type Manufacturer struct {
	WMI  string `json:"wmi" db:"wmi"`
	Name string `json:"name" db:"name"`
}
