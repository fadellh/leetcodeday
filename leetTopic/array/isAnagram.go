package array

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapS := map[rune]int{}

	for _, v := range s {
		mapS[v] = mapS[v] + 1
	}

	for _, v := range t {
		_, exist := mapS[v]

		if exist && mapS[v] > 0 {
			mapS[v] = mapS[v] - 1
		}
		if exist && mapS[v] == 0 {
			delete(mapS, v)
		}
	}

	return len(mapS) == 0
}
