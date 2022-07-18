package LeetCode

func checkIsland(grid *[][]byte, m *[][]bool, r, c int) {
	(*m)[r][c] = true
	if r != 0 {
		if (*grid)[r-1][c] == '1' && (*m)[r-1][c] == false {
			checkIsland(grid, m, r-1, c)
		}
	}
	if c != 0 {
		if (*grid)[r][c-1] == '1' && (*m)[r][c-1] == false {
			checkIsland(grid, m, r, c-1)
		}
	}
	if r != len(*grid)-1 {
		if (*grid)[r+1][c] == '1' && (*m)[r+1][c] == false {
			checkIsland(grid, m, r+1, c)
		}
	}
	if c != len((*grid)[r])-1 {
		if (*grid)[r][c+1] == '1' && (*m)[r][c+1] == false {
			checkIsland(grid, m, r, c+1)
		}
	}
}
func numIslands(grid [][]byte) int {
	res := 0
	m := make([][]bool, len(grid))
	for i := range m {
		m[i] = make([]bool, len(grid[0]))
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' && m[i][j] == false {
				res++
				checkIsland(&grid, &m, i, j)
			}
		}
	}
	return res
}
