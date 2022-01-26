package slidingWindow

func LongestSubstring(s string) int {

	word := map[string]int{}
	longest := map[string]int{}
	var exist bool
	subString := ""
	point := 0
	// reset := false
	for i, v := range s {
		_, exist = word[string(v)]

		if !exist {
			word[string(v)] = 1
			longest[string(v)] = 1
			subString += string(v)
		} else {
			word = map[string]int{}
			delete(longest, string(rune(s[point])))
			point = i
		}

	}

	return len(longest)
}
