package LeetCode

func IslandArea(grid *[][]int, m *[][]bool, r, c int) int {
	(*m)[r][c] = true
	area := 1
	if r != 0 {
		if (*grid)[r-1][c] == 1 && (*m)[r-1][c] == false {
			area += IslandArea(grid, m, r-1, c)
		}
	}
	if c != 0 {
		if (*grid)[r][c-1] == 1 && (*m)[r][c-1] == false {
			area += IslandArea(grid, m, r, c-1)
		}
	}
	if r != len(*grid)-1 {
		if (*grid)[r+1][c] == 1 && (*m)[r+1][c] == false {
			area += IslandArea(grid, m, r+1, c)
		}
	}
	if c != len((*grid)[r])-1 {
		if (*grid)[r][c+1] == 1 && (*m)[r][c+1] == false {
			area += IslandArea(grid, m, r, c+1)
		}
	}
	return area
}
func maxAreaOfIsland(grid [][]int) int {
	res := 0
	a := 0
	m := make([][]bool, len(grid))
	for i := range m {
		m[i] = make([]bool, len(grid[0]))
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 && m[i][j] == false {
				a = IslandArea(&grid, &m, i, j)
				if a > res {
					res = a
				}
			}
		}
	}
	return res
}
