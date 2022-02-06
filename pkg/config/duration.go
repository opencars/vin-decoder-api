package config

import (
	"time"
)

// Duration represents custom type for unmarshalling string.
// For example: "500ms", "1s", "2m", etc.
type Duration struct {
	time.Duration
}

// UnmarshalText implements yaml unmarshaler.
func (d *Duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))

	return err
}
