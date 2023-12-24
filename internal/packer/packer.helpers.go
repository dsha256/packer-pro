package packer

import (
	"context"

	"github.com/dsha256/packer-pro/internal/entity"
	entitysize "github.com/dsha256/packer-pro/internal/entity/size"
)

// getMinNecessaryPacks calculates minimum packs quantity for given items based on packs sizes.
func getMinNecessaryPacks(items int, sortedSizes []int) map[int]int {
	necessaryPacks := make(map[int]int)
	lastUsedPackIndex := len(sortedSizes) - 1

	for lastUsedPackIndex > 0 {
		if items-sortedSizes[lastUsedPackIndex] >= 0 {
			necessaryPacks[sortedSizes[lastUsedPackIndex]]++
			items -= sortedSizes[lastUsedPackIndex]
		} else {
			if _, exists := necessaryPacks[sortedSizes[lastUsedPackIndex]]; exists {
				diff := sortedSizes[lastUsedPackIndex] - items
				isNotCompatibleWithOtherSizes := false
				for _, size := range sortedSizes[:lastUsedPackIndex] {
					if size > diff {
						isNotCompatibleWithOtherSizes = true
						break
					}
				}
				if isNotCompatibleWithOtherSizes {
					necessaryPacks[sortedSizes[lastUsedPackIndex]]++
					items -= sortedSizes[lastUsedPackIndex]
					break
				}
			}
			lastUsedPackIndex--
		}
	}

	if items > 0 {
		for _, size := range sortedSizes {
			if size >= items {
				necessaryPacks[size]++
				break
			}
		}
	}

	return necessaryPacks
}

// refreshSortedSizesCacheFromDB retrieves sorted sizes from the db and caches them.
func refreshSortedSizesCacheFromDB(ctx context.Context, entity *entity.Client) error {
	err := cleanUpSortedSizesCache(ctx)
	if err != nil {
		return err
	}

	allSizes, err := entity.Size.Query().Select(entitysize.FieldSize).All(ctx)
	if err != nil {
		return err
	}

	for _, size := range allSizes {
		_, err = SortedSizesCache.PushRight(ctx, sortedSizesKey, size.Size)
		if err != nil {
			return err
		}
	}

	return nil
}

// cleanUpSortedSizesCache cleans up the redis cluster's keyspace dedicated for SortedSizesCache.
func cleanUpSortedSizesCache(ctx context.Context) error {
	_, err := SortedSizesCache.Delete(ctx, sortedSizesKey)
	if err != nil {
		return err
	}

	return nil
}

// refreshSortedSizesCache sets new sorted sizes to SortedSizesCache.
func refreshSortedSizesCache(ctx context.Context, sortedSizes []int) error {
	for _, size := range sortedSizes {
		_, err := SortedSizesCache.PushRight(ctx, sortedSizesKey, size)
		if err != nil {
			return err
		}
	}

	return nil
}

// cleanUpDBSizes deletes all the sizes stored in the DB.
func cleanUpDBSizes(ctx context.Context, entity *entity.Client) error {
	_, err := entity.Size.Delete().Where().Exec(ctx)
	if err == nil {
		return err
	}

	return nil
}
