package packer

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
