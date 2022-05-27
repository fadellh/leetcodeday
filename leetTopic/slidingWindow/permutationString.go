package slidingWindow

import "fmt"

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
			} else if prevWord == v {
				continue
			}
			counter = 0
			continue
		}
		counter = 0
	}
	return len(s1) == counter
}

func CheckInclusion2(s1 string, s2 string) bool {
	words := map[rune]int{}

	for _, v := range s1 {
		words[v] = words[v] + 1
	}
	// count := len(s1)
	tempCount := 0
	sample := words

	for _, v := range s2 {
		_, exist := sample[v]
		if len(s1) == tempCount {
			return true
		}

		if exist {
			if sample[v] == 1 {
				sample[v] = 0
				tempCount++
				continue
			}
		}
		tempCount--
		sample = words
	}

	return len(s1) == tempCount
}

type Char struct {
	Alpha rune
	Count int
}

func CheckInclusionLeet(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1Map := [26]int{}
	s2Map := [26]int{}

	for i, v := range s1 {
		s1Map[v-97]++
		s2Map[s2[i]-97]++
	}

	diff := len(s2) - len(s1)

	for i := 0; i < diff; i++ {
		if matches(s1Map, s2Map) {
			return true
		}

		a := i + len(s1)
		abjad := s2[a]
		abjadRemove := s2[i]

		strA := string(abjad)
		strRem := string(abjadRemove)

		fmt.Println(strA)
		fmt.Println(strRem)
		iAbjad := abjad - 97
		iAbRem := abjadRemove - 97

		s2Map[iAbjad]++
		s2Map[iAbRem]--
	}
	return matches(s1Map, s2Map)
}

func matches(s1 [26]int, s2 [26]int) bool {
	for i := 0; i < 26; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
