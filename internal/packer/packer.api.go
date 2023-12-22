package packer

import (
	"context"
	"errors"
)

var (
	// ErrorInvalidItems ...
	ErrorInvalidItems = errors.New("items must be a positive number")
)

// GetPacketsReq ...
type GetPacketsReq struct {
	Items int `query:"items"`
}

// Validate validates GetPacketsReq.
func (req *GetPacketsReq) Validate() error {
	if req.Items <= 0 {
		return ErrorInvalidItems
	}

	return nil
}

// GetPacketsResp ...
type GetPacketsResp struct {
	Packets map[int]int `json:"packets"`
}

//encore:api public method=GET path=/api/v1/packets
func (packer *Packer) GetPackets(ctx context.Context, items *GetPacketsReq) (*GetPacketsResp, error) {
	var response GetPacketsResp
	ss, err := SortedSizes.Items(ctx, sortedSizesKey)
	if err != nil {
		return &response, err
	}
	response.Packets = getMinNecessaryPacks(items.Items, ss)

	return &response, nil
}
