package LeetCode

func uniquePaths(m int, n int) int {
	arr := make([][]int, m)
	for i := range arr {
		arr[i] = make([]int, n)
	}
	for i := range arr {
		for j := range arr[i] {
			if i+j < 2 {
				arr[i][j] = 1
			} else {
				if i != 0 {
					arr[i][j] += arr[i-1][j]
				}
				if j != 0 {
					arr[i][j] += arr[i][j-1]
				}
			}
		}
	}
	return arr[m-1][n-1]
}
