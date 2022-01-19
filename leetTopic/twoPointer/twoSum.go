package twopointer

func TwoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	var diff int
	for left < right {
		diff = numbers[right] + numbers[left]

		if diff == target {
			return []int{(left + 1), (right + 1)}
		} else if diff > target {
			right--
			continue
		} else {
			left++
		}
	}

	return []int{}
}
