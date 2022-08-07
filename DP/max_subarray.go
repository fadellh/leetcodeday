package dp

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	curMax := 0

	for _, v := range nums {
		sum := curMax + v

		if v >= sum {
			curMax = v
		} else {
			curMax = sum
		}

		if curMax >= maxSum {
			maxSum = curMax
		}
	}
	return maxSum
}
