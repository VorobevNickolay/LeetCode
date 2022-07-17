package LeetCode

import (
	"fmt"
	"unicode/utf8"
)

func checkInclusion(s1 string, s2 string) bool {
	eq := true
	s1Len := utf8.RuneCountInString(s1)
	s2Len := utf8.RuneCountInString(s2)
	if s1Len > s2Len {
		return false
	}
	m := make(map[rune]int)
	j := 0
	m1 := make(map[rune]int)
	for i, chr := range s1 {
		m[chr]++
		m1[rune(s2[i])]++
	}
	for _, chr := range s1 {
		if m[chr] != m1[chr] {
			eq = false
			break
		}
	}
	if eq {
		return true
	}
	for i := 0; i < s2Len-s1Len; i++ {
		eq = true
		j = i + s1Len
		m1[rune(s2[i])]--
		m1[rune(s2[j])]++
		fmt.Println(m, m1)
		fmt.Println(i, j)
		for _, chr := range s1 {
			if m[chr] != m1[chr] {
				eq = false
				break
			}
		}
		if eq {
			return true
		}
	}
	return false
}
