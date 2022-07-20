package LeetCode

func twoSum2(numbers []int, target int) []int {
	bigId := len(numbers) - 1
	littleId := 0
	for littleId != bigId {
		if numbers[littleId]+numbers[bigId] > target {
			bigId--
		} else if numbers[littleId]+numbers[bigId] < target {
			littleId++
		} else {
			return []int{littleId + 1, bigId + 1}
		}
	}
	return []int{0}
}
