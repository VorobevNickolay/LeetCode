package LeetCode

func longestPalindrome(s string) int {
	m := make(map[rune]int)
	isMiddle := false
	res := 0
	for _, chr := range s {
		m[chr]++
	}
	for _, nums := range m {
		if !isMiddle && nums%2 == 1 {
			res += 1
			isMiddle = true
		}
		res += nums / 2 * 2
	}
	return res
}
