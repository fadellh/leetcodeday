package array

func matrixReshape(mat [][]int, r int, c int) [][]int {

	// 1. Check current column
	rows := len(mat)
	column := len(mat[0])

	//2. check
	if (rows * column) != (r * c) {
		return mat
	}

	out_mat := make([][]int, r)
	for i := range out_mat {
		out_mat[i] = make([]int, c)
	}

	row_mat := 0
	column_mat := 0
	//3. Traversel in 2D array loop in loop
	for i := 0; i < rows; i++ {
		for j := 0; j < column; j++ {
			// m = append(m,mat[i][j])
			out_mat[row_mat][column_mat] = mat[i][j]
			column_mat++
			if column_mat == c {
				column_mat = 0
				row_mat++
			}
		}
	}

	return out_mat
}
