package twopointer_test

import (
	twopointer "leetcode/twoPointer"
	"testing"
)

var testArr = []struct {
	input    []int
	target   int
	expected []int
}{
	{[]int{4, 5}, 9, []int{1, 2}},
	{[]int{4, 5}, 9, []int{1, 2}},
	{[]int{4, 5}, 9, []int{1, 2}},
}

func TestTwoSum(t *testing.T) {
	t.Run("Expected right answer", func(t *testing.T) {
		for _, test := range testArr {
			outArr := twopointer.TwoSum(test.input, test.target)

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
