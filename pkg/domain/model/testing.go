package model

import "testing"

// TestManufacturer returns valid manufacturer entity for testing.
func TestManufacturer(t *testing.T) *Manufacturer {
	t.Helper()

	return &Manufacturer{
		WMI:  "5YJ",
		Name: "Tesla Inc.",
	}
}
