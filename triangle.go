package LeetCode

import "math"

func minimumTotal(triangle [][]int) int {
	res := math.MaxInt
	for i := 1; i < len(triangle); i++ {
		triangle[i][0] += triangle[i-1][0]
		triangle[i][len(triangle[i])-1] += triangle[i-1][len(triangle[i])-2]
	}
	for i := 1; i < len(triangle); i++ {
		for j := 1; j < len(triangle[i])-1; j++ {
			if triangle[i-1][j] <= triangle[i-1][j-1] {
				triangle[i][j] += triangle[i-1][j]
			} else {
				triangle[i][j] += triangle[i-1][j-1]
			}
		}
	}
	for j := 0; j < len(triangle[len(triangle)-1]); j++ {
		if res > triangle[len(triangle)-1][j] {
			res = triangle[len(triangle)-1][j]
		}
	}
	return res
}
