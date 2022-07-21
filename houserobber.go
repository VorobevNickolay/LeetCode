package LeetCode

func rob(nums []int) int {
	res := 0
	max := make([]int, len(nums))
	if len(nums) == 1 {
		return nums[0]
	}
	max[0] = nums[0]
	max[1] = nums[1]
	if len(nums) == 2 {
		if max[0] > max[1] {
			return max[0]
		} else {
			return max[1]
		}
	}
	max[2] = nums[2] + nums[0]
	for i := 3; i < len(nums); i++ {
		if max[i-3] > max[i-2] {
			max[i] = nums[i] + max[i-3]
		} else {
			max[i] = nums[i] + max[i-2]
		}
	}
	for _, num := range max {
		if num > res {
			res = num
		}
	}
	return res
}
