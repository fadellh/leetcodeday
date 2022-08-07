package twopointer

//Given an integer array nums
//sorted in non-decreasing order,
//return an array of the squares of each number sorted in non-decreasing order.
func SortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	idx := len(nums) - 1
	sqr := make([]int, len(nums))
	// sqr := []int{}

	for left <= right {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			sqr[idx] = nums[left] * nums[left]
			idx--
			left++
		} else {
			sqr[idx] = nums[right] * nums[right]
			idx--
			right--
		}
	}

	return sqr
}

// func quickSort(arr []int, start, end int) {
// 	if start >= end {
// 		return
// 	}
// 	//Partitiion
// 	boundary := partition(arr, start, end)
// 	//Sort Left
// 	quickSort(arr, start, boundary-1)
// 	//Sort Right
// 	quickSort(arr, boundary+1, end)

// }

// func partition(arr []int, start, end int) int {
// 	pivot := arr[end]
// 	boundary := start - 1

// 	for i := start; i <= end; i++ {
// 		if arr[i] <= pivot {
// 			boundary++
// 			swap(arr, i, boundary)
// 		}
// 	}

// 	return boundary
// }

// func swap(arr []int, idx1, idx2 int) {
// 	temp := arr[idx1]
// 	arr[idx1] = arr[idx2]
// 	arr[idx2] = temp
// }
