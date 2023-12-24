package packer

import (
	"context"
	"errors"
	"slices"

	"encore.dev/storage/cache"
	"github.com/dsha256/packer-pro/internal/entity"
)

var (
	// ErrorNegativeItems ...
	ErrorNegativeItems = errors.New("items must be a positive number")
	// ErrDuplicatedSizes ..
	ErrDuplicatedSizes = errors.New("sizes must be uniq")
	// ErrEmptySizes ...
	ErrEmptySizes = errors.New("sizes must contain 1 or more elements")
)

// GetPacketsReq ...
type GetPacketsReq struct {
	Items int `query:"items"`
}

// Validate validates GetPacketsReq.
func (req *GetPacketsReq) Validate() error {
	if req.Items <= 0 {
		return ErrorNegativeItems
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

	key := CalculatedSizesKey{Items: items.Items}
	calculatedPacks, err := CalculatedPacksCache.Get(ctx, key)
	if err != nil && !errors.Is(err, cache.Miss) {
		return &response, err
	}

	if errors.Is(err, cache.Miss) {
		ss, err := SortedSizesCache.Items(ctx, sortedSizesKey)
		if err != nil {
			return &response, err
		}
		necessaryPacks := getMinNecessaryPacks(items.Items, ss)
		response.Packets = necessaryPacks

		err = CalculatedPacksCache.Set(ctx, key, GetPacketsResp{Packets: necessaryPacks})
		if err != nil {
			return &response, err
		}

		return &response, err
	}
	response.Packets = calculatedPacks.Packets
	return &response, nil
}

// ListSizesResp ...
type ListSizesResp struct {
	SortedSizes []int `query:"sorted_sizes"`
}

//encore:api public method=GET path=/api/v1/packets/sizes
func (packer *Packer) ListSizes(ctx context.Context) (*ListSizesResp, error) {
	var response ListSizesResp
	ss, err := SortedSizesCache.Items(ctx, sortedSizesKey)
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

		ss, err = SortedSizesCache.Items(ctx, sortedSizesKey)
		if err != nil {
			return &response, err
		}
		response.SortedSizes = ss

		return &response, nil
	}

	response.SortedSizes = ss

	return &response, nil
}

// PostSizesReq ...
type PostSizesReq struct {
	Sizes []int `json:"sizes"`
}

// Validate validates PostSizesReq.
func (req *PostSizesReq) Validate() error {
	if len(req.Sizes) < 1 {
		return ErrEmptySizes
	}

	weights := make(map[int]int)
	for _, size := range req.Sizes {
		if _, exists := weights[size]; exists {
			return ErrDuplicatedSizes
		}
		weights[size]++
	}

	return nil
}

// PostSizesResp ...
type PostSizesResp struct {
	SortedSizes []int `json:"sorted_sizes"`
}

//encore:api public method=POST path=/api/v1/packets/sizes
func (packer *Packer) PostSizes(ctx context.Context, sizes *PostSizesReq) (*PostSizesResp, error) {
	var response PostSizesResp

	slices.Sort(sizes.Sizes)

	err := cleanUpDBSizes(ctx, packer.entity)
	if err != nil {
		return &response, err
	}

	sizeCreates := make([]*entity.SizeCreate, len(sizes.Sizes))
	for i, size := range sizes.Sizes {
		sizeCreates[i] = packer.entity.Size.Create().SetSize(size)
	}
	newSizes, err := packer.entity.Size.CreateBulk(sizeCreates...).Save(ctx)
	if err != nil {
		return &response, err
	}

	err = cleanUpSortedSizesCache(ctx)
	if err != nil {
		return &response, err
	}

	var sortedSizes []int
	for _, size := range newSizes {
		sortedSizes = append(sortedSizes, size.Size)
	}

	err = refreshSortedSizesCache(ctx, response.SortedSizes)
	if err != nil {
		return &response, err
	}

	response.SortedSizes = sortedSizes

	return &response, nil
}
