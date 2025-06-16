package rahnama

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) string {

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

	maxString := ""
	for k, _ := range uniqueSubstrings {
		viwCount := 0
		for _, b := range k {
			if b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' {
				viwCount++
			}
		}
		if viwCount%2 == 0 && viwCount != 0 && len(k) > len(maxString) {
			maxString = k
		}
	}

	return maxString
}
