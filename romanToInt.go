package LeetCode

import "unicode/utf8"

func romanToInt(s string) int {
	res := 0
	var chrNext byte
	length := utf8.RuneCountInString(s)
	for idx, chr := range s {
		if chr == 'V' {
			res += 5
		} else if chr == 'L' {
			res += 50
		} else if chr == 'D' {
			res += 500
		} else if chr == 'M' {
			res += 1000
		} else if chr == 'I' {
			if idx+1 < length {
				chrNext = s[idx+1]
			} else {
				chrNext = ' '
			}
			if chrNext == 'V' || chrNext == 'X' {
				res--
			} else {
				res++
			}
		} else if chr == 'X' {
			if idx+1 < length {
				chrNext = s[idx+1]
			} else {
				chrNext = ' '
			}
			if chrNext == 'L' || chrNext == 'C' {
				res -= 10
			} else {
				res += 10
			}
		} else if chr == 'C' {
			if idx+1 < length {
				chrNext = s[idx+1]
			} else {
				chrNext = ' '
			}
			if chrNext == 'D' || chrNext == 'M' {
				res -= 100
			} else {
				res += 100
			}
		}
	}
	return res
}
