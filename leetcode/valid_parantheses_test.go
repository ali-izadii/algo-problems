package leetcode

import "testing"

func TestValidParentheses(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    bool
	}{
		{description: "Test1", input: "{{{}", expected: false},
		{description: "Test2", input: "()[]{}", expected: true},
		{description: "Test3", input: "(", expected: false},
		{description: "Test4", input: "))", expected: false},
		{description: "Test5", input: "([])", expected: true},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := isValid(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %t, got %t", tc.expected, result)
			}
		})
	}
}
