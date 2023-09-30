package order

import (
	"github.com/anvlad11/testapp-20230927/pkg/config"
	"sort"
)

type Service struct {
	packSizes    []int
	packSizesMap map[int]interface{}
}

func NewService(cfg config.Orders) *Service {
	packSizes := []int{}
	packSizesMap := map[int]interface{}{}
	for _, size := range cfg.PackSizes {
		if _, exists := packSizesMap[size]; !exists {
			packSizesMap[size] = nil
			packSizes = append(packSizes, size)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	return &Service{
		packSizes:    packSizes,
		packSizesMap: packSizesMap,
	}
}
