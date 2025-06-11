package hacker_rank

import (
	"strings"
)

// https://www.hackerrank.com/challenges/two-characters/problem?isFullScreen=true

func TwoCharacters(s string) int32 {
	charSet := make(map[rune]bool)
	for _, char := range s {
		charSet[char] = true
	}

	var unique []rune
	for char := range charSet {
		unique = append(unique, char)
	}

	maxLen := 0
	for i := 0; i < len(unique); i++ {
		for j := i + 1; j < len(unique); j++ {
			c1, c2 := unique[i], unique[j]

			var sb strings.Builder
			for _, r := range s {
				if r == c1 || r == c2 {
					sb.WriteRune(r)
				}
			}

			filteredString := sb.String()
			flag := true
			for i := 0; i < len(filteredString)-1; i++ {
				if filteredString[i] == filteredString[i+1] {
					flag = false
					break
				}
			}

			if flag {
				maxLen = max(maxLen, len(filteredString))
			}
		}
	}
	return int32(maxLen)
}
