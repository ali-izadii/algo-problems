package hacker_rank

import "testing"

func TestTwoCharacters(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    int32
	}{
		{description: "Test1", input: "beabeefeab", expected: 5},
		{description: "Test2", input: "aab", expected: 0},
		{description: "Test3", input: "aaa", expected: 0},
		{description: "Test4", input: "abcdef", expected: 2},
		{description: "Test5", input: "", expected: 0},
		{description: "Test6", input: "a", expected: 0},
		{description: "Test7", input: "ab", expected: 2},
		{description: "Test8", input: "abab", expected: 4},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := TwoCharacters(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}
