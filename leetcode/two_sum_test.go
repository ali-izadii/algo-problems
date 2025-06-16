package leetcode

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		description string
		input       []int
		target      int
		expected    []int
	}{
		{description: "Test2", input: []int{3, 3}, target: 6, expected: []int{0, 1}},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := twoSum(tc.input, tc.target)
			fmt.Println(result)
			//if result != tc.expected {
			//	t.Errorf("Expected %s, got %s", tc.expected, result)
			//}
		})
	}
}
