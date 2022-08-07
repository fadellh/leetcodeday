package binarysearch

func SearchInsert(nums []int, target int) int {

	left := 0
	right := len(nums) - 1
	mid := 0

	for left <= right {
		mid = (right + left) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
			continue
		}
		right = mid - 1

	}
	if target > nums[mid] {
		return mid + 1
	} else {
		return mid
	}
}
