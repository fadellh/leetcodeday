package twopointer

func ReverseString(s []byte) []byte {
	left := 0
	right := len(s) - 1

	var temp byte
	for left <= right {
		//swap
		temp = s[left]
		s[left] = s[right]
		s[right] = temp
		left++
		right--
	}

	return s
}
