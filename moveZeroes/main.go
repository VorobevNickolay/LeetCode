package main

import "fmt"

func moveZeroes(nums []int) {
	nulls := 0
	for i, n := range nums {
		if n == 0 {
			nums[i] = nums[len(nums)-i-1]
			nulls++
		}
	}
}
func main() {
	arr := []int{0, 1, 0, 3, 1}
	moveZeroes(arr)
	fmt.Print(arr)
}
