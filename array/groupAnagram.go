package array

type Count struct {
	Arr [26]int
}

func groupAnagrams(strs []string) [][]string {
	group := map[[26]int][]string{}

	for _, v := range strs {
		count := [26]int{}

		for _, c := range v {
			count[c-97] = count[c-97] + 1
		}

		_, exist := group[count]
		if exist {
			val := group[count]
			val = append(val, v)
			group[count] = val
			continue
		}
		group[count] = []string{v}
	}
	result := [][]string{}
	for _, v := range group {
		result = append(result, v)
	}

	return result
}
