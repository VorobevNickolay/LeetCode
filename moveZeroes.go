package LeetCode

func moveZeroes(nums []int) {
	nulls := 0
	for i, n := range nums {
		if n == 0 {
			nums[i] = nums[len(nums)-i-1]
			nulls++
		}
	}
}
