package calculator

import (
	"reflect"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	tests := []struct {
		order     int
		packSizes []int
		expected  map[int]int
	}{
		{
			order:     500000,
			packSizes: []int{23, 31, 53},
			expected:  map[int]int{23: 2, 31: 7, 53: 9429},
		},
		{
			order:     14,
			packSizes: []int{12, 5},
			expected:  map[int]int{5: 3},
		},
		{
			order:     1,
			packSizes: []int{1000, 500, 250},
			expected:  map[int]int{250: 1},
		},
		{
			order:     0,
			packSizes: []int{5, 10, 15},
			expected:  map[int]int{},
		},
	}

	for _, test := range tests {
		result := CalculatePacks(test.order, test.packSizes)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For order %d and pack sizes %v, expected %v but got %v",
				test.order, test.packSizes, test.expected, result)
		}
	}
}
