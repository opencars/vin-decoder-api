package command

import (
	"regexp"
	"strings"
	"time"

	"github.com/opencars/schema"
	"github.com/opencars/schema/vehicle"
	"github.com/opencars/seedwork"
	"google.golang.org/protobuf/types/known/timestamppb"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	IsVIN = regexp.MustCompile(`^[A-HJ-NPR-Z0-9]{17}$`)
)

type DecodeVIN struct {
	UserID  string
	TokenID string
	VIN     string
}

func (c *DecodeVIN) Prepare() {
	c.VIN = strings.ReplaceAll(strings.ToUpper(c.VIN), "-", "")
}

func (c *DecodeVIN) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(
			&c.UserID,
			validation.Required.Error(seedwork.Required),
		),
		validation.Field(
			&c.TokenID,
			validation.Required.Error(seedwork.Required),
		),
		validation.Field(
			&c.VIN,
			validation.Required.Error(seedwork.Required),
			validation.Match(IsVIN).Error(seedwork.Invalid),
		),
	)
}

func (c *DecodeVIN) Event() schema.Producable {
	msg := vehicle.VINDecoded{
		UserId:     c.UserID,
		TokenId:    c.TokenID,
		Vin:        c.VIN,
		SearchedAt: timestamppb.New(time.Now().UTC()),
	}

	return schema.New(&source, &msg).Message(
		schema.WithSubject(schema.VinDecodingCustomerActions),
	)
}
