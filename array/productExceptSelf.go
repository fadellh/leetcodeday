package array

func productExceptSelf2(nums []int) []int {

	temp := []int{}
	answer := []int{}
	for i, v := range nums {
		if len(temp) == 0 {
			temp = append(temp, v)
			total := 1
			for j, m := range nums {
				if j == i {
					continue
				}
				total *= m
			}
			answer = append(answer, total)
			continue
		}

		total := 1
		for j, m := range nums {
			if i == j {
				continue
			}
			if len(temp)-1 >= j {
				total *= temp[j]
				continue
			}
			total *= m
		}
		answer = append(answer, total)
		temp = append(temp, v)
	}

	return answer
}

func productExceptSelf(nums []int) []int {
	answer := []int{}

	prefix := 1
	for _, v := range nums {
		answer = append(answer, prefix)
		prefix *= v
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] = answer[i] * postfix
		postfix *= nums[i]
	}
	return answer
}
