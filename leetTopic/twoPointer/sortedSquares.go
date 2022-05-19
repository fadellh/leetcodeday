package twopointer

func SortedSquares(nums []int) []int {
	sqr := []int{}

	for _, v := range nums {
		sqr = append(sqr, v*v)
	}
	quickSort(sqr, 0, len(nums)-1)

	return sqr
}

func quickSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	//Partitiion
	boundary := partition(arr, start, end)
	//Sort Left
	quickSort(arr, start, boundary-1)
	//Sort Right
	quickSort(arr, boundary+1, end)

}

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	boundary := start - 1

	for i := start; i <= end; i++ {
		if arr[i] <= pivot {
			boundary++
			swap(arr, i, boundary)
		}
	}

	return boundary
}

func swap(arr []int, idx1, idx2 int) {
	temp := arr[idx1]
	arr[idx1] = arr[idx2]
	arr[idx2] = temp
}
