package LeetCode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if _, ok := m[num]; ok {
			return []int{m[num], i}
		}
		m[target-num] = i
	}
	return []int{0, 1}
}
