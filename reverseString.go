package LeetCode

import "strings"

func reverseString(s []byte) string {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return string(s)
}
func reverseWords(s string) string {
	res := ""
	var b []byte
	str := strings.Split(s, " ")
	for i := range str {
		b = []byte(str[i])
		if i != 0 {
			res += " "
		}
		res += reverseString(b)
	}
	return res
}
