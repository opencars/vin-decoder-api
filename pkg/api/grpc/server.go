package grpc

import (
	"context"

	"github.com/opencars/grpc/pkg/vin_decoding"

	"github.com/opencars/vin-decoder-api/pkg/domain/command"
)

type vinDecodingHandler struct {
	vin_decoding.UnimplementedServiceServer
	api *API
}

func (h *vinDecodingHandler) Decode(ctx context.Context, r *vin_decoding.DecodeRequest) (*vin_decoding.DecodeResultList, error) {
	c := command.DecodeVINInternal{
		Items: make([]command.Item, 0, len(r.Vins)),
	}

	for _, vin := range r.Vins {
		c.Items = append(c.Items, command.Item{VIN: vin})
	}

	result, err := h.api.svc.Decode(ctx, &c)
	if err != nil {
		return nil, handleErr(err)
	}

	dto := vin_decoding.DecodeResultList{
		Items: make([]*vin_decoding.DecodeResultItem, 0, len(result.Results)),
	}

	for i := range result.Results {
		dto.Items = append(dto.Items, ResultItemFromDomain(&result.Results[i]))
	}

	return &dto, nil
}
