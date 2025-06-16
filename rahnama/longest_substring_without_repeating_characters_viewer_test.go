package rahnama

import "testing"

func TestTLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    string
	}{
		{description: "Test2", input: "abcciiidef", expected: "idef"},
		{description: "Test3", input: "abcabcbb", expected: ""},
		{description: "Test4", input: "aeiouxyz", expected: "eiouxyz"},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := lengthOfLongestSubstring(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
