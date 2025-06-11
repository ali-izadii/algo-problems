package leetcode

import "fmt"

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {

	uniqueSubstrings := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		seen := make(map[byte]bool)
		for j := i; j < len(s); j++ {
			if seen[s[j]] {
				break
			}
			seen[s[j]] = true
			uniqueSubstrings[s[i:j+1]] = true
		}
	}

	for k, _ := range uniqueSubstrings {
		fmt.Println(k)
	}

	return 0
}
