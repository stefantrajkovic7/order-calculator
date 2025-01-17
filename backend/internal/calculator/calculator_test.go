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
			order:     501,
			packSizes: []int{250, 500, 1000},
			expected:  map[int]int{500: 1, 250: 1},
		},
		{
			order:     12001,
			packSizes: []int{250, 500, 1000, 2000, 5000},
			expected:  map[int]int{5000: 2, 2000: 1, 250: 1},
		},
		{
			order:     1,
			packSizes: []int{250, 500, 1000},
			expected:  map[int]int{250: 1},
		},
	}

	for _, test := range tests {
		result := CalculatePacks(test.order, test.packSizes)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For order %d and pack sizes %v, expected %v but got %v", test.order, test.packSizes, test.expected, result)
		}
	}
}
