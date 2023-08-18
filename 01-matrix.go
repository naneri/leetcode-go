// https://leetcode.com/problems/01-matrix

type Cell struct {
	row, col int
}

func updateMatrix(mat [][]int) [][]int {

	m, n := len(mat), len(mat[0])
	queue := make([]Cell, 0)
	MAX_VALUE := m * n

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, Cell{i, j})
			} else {
				mat[i][j] = MAX_VALUE
			}
		}
	}

	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		for _, dir := range directions {
			r, c := cell.row+dir[0], cell.col+dir[1]
			if inBounds(r, c, m, n) && mat[r][c] > matrixCellValue(mat, cell)+1 {
				queue = append(queue, Cell{r, c})
				mat[r][c] = matrixCellValue(mat, cell) + 1
			}
		}
	}
	
	return mat
}

func matrixCellValue(mat [][]int, cell Cell) int {
	return mat[cell.row][cell.col]	
}

func inBounds(r, c, m, n int) bool {
	return r >= 0 && r < m && c >= 0 && c < n
}
