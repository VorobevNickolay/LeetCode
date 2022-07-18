package LeetCode

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)

	res := 0
	cur := 0
	start := 0

	for idx, chr := range s {
		j, ok := m[chr]
		if !ok || j < idx-cur {
			cur++
		} else {
			if cur > res {
				res = cur
			}
			start = j + 1
			cur = idx - start + 1
		}
		m[chr] = idx
	}
	if cur > res {
		res = cur
	}
	return res
}
