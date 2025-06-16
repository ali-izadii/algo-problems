package leetcode

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	t := true
	l := len(s)
	for i := 0; i < l/2; i++ {
		t = s[i] == s[l-1-i] && t
	}

	return t
}
