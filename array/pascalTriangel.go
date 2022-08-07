package array

func generate(numRows int) [][]int {
	m := [][]int{{1}}

	if numRows == 1 {
		return m
	}

	m = append(m, []int{1, 1})
	if numRows == 2 {
		return m
	}

	m = append(m, []int{1, 2, 1})
	sum := 0

	for i := 2; i < numRows-1; i++ {
		c := []int{}
		left := -1
		right := 0
		for left <= len(m[i])-1 {
			if left == len(m[i])-1 {
				sum = m[i][right-1] + 0
				c = append(c, sum)
				left++
				continue
			}
			if left == -1 {
				sum = 0 + m[i][right]
				c = append(c, sum)
				left++
				right++
			}
			sum = m[i][left] + m[i][right]
			c = append(c, sum)
			left++
			right++
		}
		m = append(m, c)
	}

	return m

}
