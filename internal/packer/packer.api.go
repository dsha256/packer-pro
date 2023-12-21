package packer

import (
	"context"
	_ "embed"
	"errors"
)

var (
	ErrorInvalidItems = errors.New("items must be a positive number")
)

// GetPacketsReq ...
type GetPacketsReq struct {
	Items int `json:"items"`
}

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
func GetPackets(ctx context.Context, items *GetPacketsReq) (*GetPacketsResp, error) {
	return &GetPacketsResp{}, nil
}
