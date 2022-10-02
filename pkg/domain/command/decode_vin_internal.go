package command

import (
	"strings"
)

type DecodeVINInternal struct {
	VINs []string
}

func (c *DecodeVINInternal) Prepare() {
	for i, vin := range c.VINs {
		c.VINs[i] = strings.ReplaceAll(strings.ToUpper(vin), "-", "")
	}
}

func (c *DecodeVINInternal) Validate() error {
	return nil
}
