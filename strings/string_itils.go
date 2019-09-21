package strings

import (
	"strings"
)

func IsPalindrome(s string) bool {
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func IsPalindromeIgnoreCase(s string) bool {
	return IsPalindrome(strings.ToLower(s))
}
