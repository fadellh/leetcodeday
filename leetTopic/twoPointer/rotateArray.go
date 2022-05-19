package twopointer

func Rotate(nums []int, k int) {
	arr := make([]int, len(nums))
	var newIdx int

	for i, v := range nums {
		newIdx = i + k

		if newIdx >= len(nums)-1 {
			mod := newIdx % len(nums)
			newIdx = mod
		}

		arr[newIdx] = v
	}
	copy(nums, arr)
}
