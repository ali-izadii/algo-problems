package leetcode

// https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	index := 0
	var arr = make([]rune, len(s))
	for _, n := range s {
		if n == '(' || n == '{' || n == '[' {
			arr[index] = n
			index = index + 1
		} else {
			if index == 0 {
				return false
			}

			if n == ')' && arr[index-1] == '(' {
				index = index - 1
			} else if n == '}' && arr[index-1] == '{' {
				index = index - 1
			} else if n == ']' && arr[index-1] == '[' {
				index = index - 1
			} else {
				return false
			}
		}
	}
	return index == 0
}
