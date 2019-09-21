package strings_test

import (
	"github.com/stretchr/testify/assert"
	"go-sandbox/strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, strings.IsPalindrome("mama"), false, "Not palindrome")
	assert.Equal(t, strings.IsPalindrome("abbaGabba"), true, "Palindrome")
	assert.Equal(t, strings.IsPalindrome("Several words sdrow lareveS"), true, "Palindrome")
}

func TestIsPalindromeIgnoreCase(t *testing.T) {
	assert.Equal(t, strings.IsPalindromeIgnoreCase("mama"), false, "Not palindrome")
	assert.Equal(t, strings.IsPalindromeIgnoreCase("AbbaGabBa"), true, "Palindrome")
	assert.Equal(t, strings.IsPalindromeIgnoreCase("Several words sdrow lareves"), true, "Palindrome")
}