package array

import "fmt"

// It's interesting. It's like mark coordinat for each number
func isValidSudoku(board [][]byte) bool {
	mapRow := map[string]bool{}
	mapCol := map[string]bool{}
	mapSqr := map[string]bool{}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == 46 {
				continue
			}
			row := fmt.Sprintf("%d-%s", r, string(board[r][c]))
			col := fmt.Sprintf("%d-%s", c, string(board[r][c]))
			sqr := fmt.Sprintf("%d-%d-%s", r/3, c/3, string(board[r][c]))
			if mapRow[row] || mapCol[col] || mapSqr[sqr] {
				return false
			}
			mapRow[row] = true
			mapCol[col] = true
			mapSqr[sqr] = true
		}
	}
	return true
}
