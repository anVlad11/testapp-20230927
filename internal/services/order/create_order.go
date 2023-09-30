package order

import (
	"github.com/anvlad11/testapp-20230927/pkg/errors"
	"math"
)

func (s *Service) CreateOrder(itemsCount int) (map[int]int, error) {
	if itemsCount <= 0 {
		return nil, errors.ItemCountIsEqualOrLessThanZero
	}

	if _, exists := s.packSizesMap[itemsCount]; exists {
		return map[int]int{itemsCount: 1}, nil
	}

	minItemsOversell := math.MaxInt
	var optimalPackCombination []int

	// Helper function to generate all possible combinations of packing
	var generateCombinations func(int, []int)
	generateCombinations = func(idx int, counts []int) {
		if idx == len(s.packSizes) {
			totalValue := 0
			for i, count := range counts {
				totalValue += s.packSizes[i] * count
			}

			if totalValue < itemsCount {
				return // Skip if totalValue is less than the required item amount
			}

			oversell := totalValue - itemsCount
			// Update if the current combination is more optimal than the previously found ones
			if oversell < minItemsOversell || (oversell == minItemsOversell && sum(counts) < sum(optimalPackCombination)) {
				minItemsOversell = oversell
				optimalPackCombination = append([]int(nil), counts...)
			}
			return
		}

		for count := 0; count <= itemsCount/s.packSizes[idx]+1; count++ {
			generateCombinations(idx+1, append(counts, count))
		}
	}

	generateCombinations(0, []int{})

	// Construct the combination of packs from the counts
	result := map[int]int{}
	for i, count := range optimalPackCombination {
		for j := 0; j < count; j++ {
			result[s.packSizes[i]]++
		}
	}

	return result, nil
}

// Helper function to calculate the sum of a slice
func sum(arr []int) int {
	s := 0
	for _, v := range arr {
		s += v
	}
	return s
}
