package LeetCode

func checkField(image *[][]int, m *[][]bool, r, c int) {
	if r != 0 {
		if (*image)[r-1][c] == (*image)[r][c] {
			if !(*m)[r-1][c] {

				(*m)[r-1][c] = true
				checkField(image, m, r-1, c)
			}
		}
	}
	if c != 0 {
		if (*image)[r][c-1] == (*image)[r][c] {
			if !(*m)[r][c-1] {
				(*m)[r][c-1] = true
				checkField(image, m, r, c-1)
			}
		}
	}
	if r != len(*image)-1 {

		if (*image)[r+1][c] == (*image)[r][c] {
			if !(*m)[r+1][c] {
				(*m)[r+1][c] = true
				checkField(image, m, r+1, c)
			}
		}
	}
	if c != len((*image)[r])-1 {
		if (*image)[r][c+1] == (*image)[r][c] {
			if !(*m)[r][c+1] {
				(*m)[r][c+1] = true
				checkField(image, m, r, c+1)
			}
		}
	}
}
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == color {
		return image
	}
	m := make([][]bool, len(image))
	for i := range m {
		m[i] = make([]bool, len(image[0]))
	}
	m[sr][sc] = true
	checkField(&image, &m, sr, sc)
	for i := range image {
		for j := range image[i] {
			if m[i][j] {
				image[i][j] = color
			}
		}
	}
	return image
}
