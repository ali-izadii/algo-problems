package leetcode

import "testing"

func TestTLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    string
	}{
		{description: "Test1", input: "abcaafhb", expected: ""},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			lengthOfLongestSubstring(tc.input)
			//if result != tc.expected {
			//	t.Errorf("Expected %d, got %d", tc.expected, result)
			//}
		})
	}
}
