package LeetCode

import "strconv"

func getHint(secret string, guess string) string {
	bulls := 0
	cows := 0
	var min int
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	for i, chr := range secret {
		if chr == rune(guess[i]) {
			bulls++
		} else {
			m1[chr]++
			m2[rune(guess[i])]++
		}
	}
	for idx := range m1 {
		if m1[idx] <= m2[idx] {
			min = m1[idx]
		} else {
			min = m2[idx]
		}
		cows += min
	}
	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}
