package leetcode

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		target      int
		expected    []int
	}{
		{description: "Test1", input: []int{3, 3}, target: 6, expected: []int{0, 1}},
		{description: "Test2", input: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := twoSum(tc.input, tc.target)
			if !slices.Equal(tc.expected, result) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
