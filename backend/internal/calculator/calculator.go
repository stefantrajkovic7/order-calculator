package calculator

import (
	"math"
	"sort"
)

// CalculatePacks determines the optimal distribution of packs for the given order.
func CalculatePacks(order int, packSizes []int) map[int]int {
	if order <= 0 || len(packSizes) == 0 {
		return map[int]int{}
	}

	// Sort pack sizes in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))

	// Limit DP table size for very large orders
	maxOrder := min(order, 100000)
	dp := make([]int, maxOrder+1)
	packTracker := make([]map[int]int, maxOrder+1)

	for i := 0; i <= maxOrder; i++ {
		dp[i] = math.MaxInt32
		packTracker[i] = make(map[int]int)
	}
	dp[0] = 0

	// Build DP table
	for _, size := range packSizes {
		for currentOrder := size; currentOrder <= maxOrder; currentOrder++ {
			if dp[currentOrder-size]+1 < dp[currentOrder] {
				dp[currentOrder] = dp[currentOrder-size] + 1
				packTracker[currentOrder] = copyMap(packTracker[currentOrder-size])
				packTracker[currentOrder][size]++
			}
		}
	}

	// If the exact order can be fulfilled within the DP table
	if order <= maxOrder && dp[order] != math.MaxInt32 {
		return packTracker[order]
	}

	// Fallback to greedy solution for large orders
	return applyGreedyFallback(order, packSizes)
}

func applyGreedyFallback(order int, packSizes []int) map[int]int {
	result := make(map[int]int)
	remainingOrder := order

	// Sort pack sizes in ascending order for better granularity
	sort.Ints(packSizes)

	// Use packs to cover the order
	for _, size := range packSizes {
		if remainingOrder == 0 {
			break
		}
		count := remainingOrder / size
		if count > 0 {
			result[size] = count
			remainingOrder -= count * size
		}
	}

	// If there's leftover order that can't be fulfilled exactly
	if remainingOrder > 0 {
		smallestPack := packSizes[0]
		result[smallestPack] += 1
	}

	return refineFallback(order, result, packSizes)
}

func refineFallback(order int, initialResult map[int]int, packSizes []int) map[int]int {
	dp := make([]int, order+1)
	combination := make([]map[int]int, order+1)

	for i := range dp {
		dp[i] = math.MaxInt32
		combination[i] = make(map[int]int)
	}
	dp[0] = 0

	// Build DP table
	for _, size := range packSizes {
		for currentOrder := size; currentOrder <= order; currentOrder++ {
			if dp[currentOrder-size]+1 < dp[currentOrder] {
				dp[currentOrder] = dp[currentOrder-size] + 1
				combination[currentOrder] = copyMap(combination[currentOrder-size])
				combination[currentOrder][size]++
			}
		}
	}

	// Use DP result if it exists
	if dp[order] != math.MaxInt32 {
		return combination[order]
	}

	return initialResult
}

func copyMap(original map[int]int) map[int]int {
	copied := make(map[int]int)
	for k, v := range original {
		copied[k] = v
	}
	return copied
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
