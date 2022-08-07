package array

func searchMatrix(matrix [][]int, target int) bool {
	mid := 0
	for i, v := range matrix {
		if target == v[0] {
			return true
		}
		if target < v[0] {
			mid = i - 1
			break
		}
		if target > v[0] {
			mid = i
		}

	}
	if mid < 0 {
		return false
	}

	ml := 0
	mr := len(matrix[mid]) - 1

	for ml <= mr {
		mm := (ml + mr) / 2
		if matrix[mid][mm] == target {
			return true
		}
		if target > matrix[mid][mm] {
			ml = mm + 1
			continue
		}
		mr = mm - 1
	}
	return false
}
