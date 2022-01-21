package twopointer

func ReverseWord(s string) string {

	var current string
	var result string
	for _, v := range s {

		if string(v) == " " {
			result = result + current + string(v)
			current = ""
			continue
		}

		current = string(v) + current

	}

	return result + current

	//Leetcode space low
	// res := make([]byte, len(s))
	// j := 0
	// for i := 0; i < len(s); i, j = i, j+1 {
	// 	for ; j < len(s) && s[j] != ' '; j += 1 {
	// 	}
	// 	for k := 0; i < j; i, k = i+1, k+1 {
	// 		res[i] = s[j-k-1]
	// 	}
	// 	if j < len(s) {
	// 		res[j] = ' '
	// 		i += 1
	// 	}
	// }
	// return string(res)

}
