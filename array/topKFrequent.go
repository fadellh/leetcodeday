package array

func topKFrequent(nums []int, k int) []int {
	hash := map[int]int{}

	for _, v := range nums {
		hash[v] = hash[v] + 1
	}

	bucket := make([][]int, len(nums))
	for key, v := range hash {
		arr := bucket[v-1]
		arr = append(arr, key)
		bucket[v-1] = arr
	}

	result := []int{}

	for i := len(bucket) - 1; i >= 0; i-- {
		if k == 0 {
			break
		}
		if len(bucket[i]) == 0 {
			continue
		}

		result = append(result, bucket[i]...)
		if len(bucket[i]) > 0 {
			k = k - len(bucket[i])
			continue
		}
		k--
	}

	return result
}
