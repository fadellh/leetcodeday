package slidingWindow

func CheckInclusion(s1 string, s2 string) bool {
	words := map[rune]int{}

	for _, v := range s1 {
		words[v] = words[v] + 1
	}

	counter := 0
	var prevWord rune

	for _, v := range s2 {
		if len(s1) == counter {
			return true
		}
		_, exist := words[v]
		if exist {
			if words[v] == 1 {
				words[v] = 0
				prevWord = v
				counter++
				continue
			} else if words[v] == 0 && prevWord != v {
				words[v] = 1
				counter++
				continue
			}
			counter = 0
			continue
		}
		counter = 0
	}
	return len(s1) == counter
}
