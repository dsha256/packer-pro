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

// ListSizesResp ...
type ListSizesResp struct {
	SortedSizes []int `query:"sorted_sizes"`
}

//encore:api public method=GET path=/api/v1/packets/sizes
func (packer *Packer) ListSizes(ctx context.Context) (*ListSizesResp, error) {
	var response ListSizesResp
	ss, err := SortedSizes.Items(ctx, sortedSizesKey)
	if err != nil {
		return &response, err
	}

	// If sorted sizes cashing time is expired, sorted sizes will be retrieved from the DB and then cashed again.
	if len(ss) == 0 {
		// making sure about cache clean up to avoid duplicated sizes in the cache
		err = cleanUpSortedSizesCache(ctx)
		if err != nil {
			return &response, err
		}

		err = refreshSortedSizesCacheFromDB(ctx, packer.entity)
		if err != nil {
			return &response, err
		}

		ss, err = SortedSizes.Items(ctx, sortedSizesKey)
		if err != nil {
			return &response, err
		}
		response.SortedSizes = ss

		return &response, nil
	}

	response.SortedSizes = ss

	return &response, nil
}
