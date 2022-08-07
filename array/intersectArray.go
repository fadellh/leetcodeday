package array

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		temp := nums1
		nums1 = nums2
		nums2 = temp
	}

	unique := map[int]int{}

	for _, v := range nums2 {
		_, exist := unique[v]
		if !exist {
			unique[v] = 1
			continue
		}
		unique[v] = unique[v] + 1
	}
	result := []int{}
	for _, v := range nums1 {
		_, exist := unique[v]

		if exist && unique[v] > 0 {
			result = append(result, v)
			unique[v] = unique[v] - 1
		}
	}

	return result

}
