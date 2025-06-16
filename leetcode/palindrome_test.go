package leetcode

import (
	"testing"
)

func TestIsPolindrom(t *testing.T) {
	testCases := []struct {
		description string
		input       int
		expected    bool
	}{
		{description: "Test1", input: 10, expected: false},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := isPalindrome(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %t, got %t", tc.expected, result)
			}
		})
	}
}
