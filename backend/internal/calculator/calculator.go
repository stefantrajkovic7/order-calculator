package calculator

import (
	"sort"
)

func CalculatePacks(order int, packSizes []int) map[int]int {
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))
	result := make(map[int]int)

	for _, size := range packSizes {
		if order == 0 {
			break
		}
		count := order / size
		if count > 0 {
			result[size] = count
			order -= count * size
		}
	}
	if order > 0 {
		result[packSizes[len(packSizes)-1]] += 1
	}
	return result
}
