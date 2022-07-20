package LeetCode

func characterReplacement(s string, k int) int {
	m := make(map[rune]int)
	max := 0
	i, j := 0, 0
	l := 0
	res := 0
	m[rune(s[0])]++
	for j < len(s) {
		l = j - i + 1
		for _, num := range m {
			if num > max {
				max = num
			}
		}
		if l-max <= k {
			if res < l {
				res = l
			}
			j++
			if j == len(s) {
				break
			}
			m[rune(s[j])]++
		} else {
			m[rune(s[i])]--
			i++
		}
	}
	return res
}
