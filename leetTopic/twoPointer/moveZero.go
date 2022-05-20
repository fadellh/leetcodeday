package twopointer

func MoveZeroes(nums []int) {
	arr := make([]int, len(nums))
	lastIdx := len(nums) - 1
	idx := 0
	for _, v := range nums {
		if v == 0 {
			arr[lastIdx] = 0
			lastIdx--
			continue
		}
		arr[idx] = v
		idx++
	}
	copy(nums, arr)
}
