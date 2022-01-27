package slidingWindow

func LongestSubstring(s string) int {

	word := map[string]int{}
	var exist bool
	longest := 0
	i := 0
	for j, v := range s {
		_, exist = word[string(v)]

		if exist {
			i = max(word[string(v)], i)
		}
		longest = max(longest, j-i+1)
		word[string(v)] = j + 1

	}

	return longest
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
