package array

func canConstruct(ransomNote string, magazine string) bool {
	mapR := map[rune]int{}
	mapM := map[rune]int{}

	for _, v := range ransomNote {
		mapR[v] = mapR[v] + 1
	}

	for _, v := range magazine {
		mapM[v] = mapM[v] + 1
	}

	for key := range mapR {
		if mapR[key] > mapM[key] {
			return false
		}
	}
	return true
}
