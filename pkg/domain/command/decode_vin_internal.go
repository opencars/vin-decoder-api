package command

import (
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/opencars/seedwork"
	"github.com/opencars/vin-decoder-api/pkg/domain/model"
)

type Item struct {
	VIN string
}

func (c *Item) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(
			&c.VIN,
			validation.Required.Error(seedwork.Required),
			validation.Match(model.IsVIN).Error(seedwork.Invalid),
		),
	)
}

type DecodeVINInternal struct {
	Items []Item
}

func (c *DecodeVINInternal) Prepare() {
	for i := range c.Items {
		c.Items[i].VIN = strings.ReplaceAll(strings.ToUpper(c.Items[i].VIN), "-", "")
	}
}

func (c *DecodeVINInternal) Validate() error {
	return nil
}
