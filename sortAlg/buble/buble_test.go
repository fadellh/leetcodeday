package buble_test

import (
	"leetcode/sortAlg/buble"
	"testing"
)

var testArr = []struct {
	input    []int
	expected []int
}{
	{[]int{4, 5}, []int{4, 5}},
	{[]int{4, 5}, []int{4, 5}},
	{[]int{4, 5}, []int{4, 5}},
}

// test function
func TestBubleSort(t *testing.T) {
	t.Run("Expected right answer", func(t *testing.T) {
		for _, test := range testArr {
			outArr := buble.BubleSort(test.input)

			if len(test.expected) != len(outArr) {
				t.Error("Test Failed. Expected: ", test.expected, " Output: ", outArr)
			}

			for idx, item := range test.expected {
				if item != outArr[idx] {
					t.Error("Test Failed. Expected: ", test.expected, " Output: ", outArr)
				}
			}

		}
	})
}
