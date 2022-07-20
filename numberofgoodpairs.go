package LeetCode

func numIdenticalPairs(nums []int) int {
	res := 0
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for i := range m {
		res += (m[i] * (m[i] - 1)) / 2
	}
	return res
}
