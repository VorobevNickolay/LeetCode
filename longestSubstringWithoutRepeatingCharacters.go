package LeetCode

import (
	"unicode/utf8"
)

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	i := 0
	res := 0
	sLen := utf8.RuneCountInString(s)
	if sLen == 0 || sLen == 1 {
		return sLen
	}
	for j := 0; j < sLen; j++ {
		if m[rune(s[j])] > i {
			i = m[rune(s[j])]
		}
		if j-i > res {
			res = j - i
		}
		m[rune(s[j])] = j

	}
	return res
}
