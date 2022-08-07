package twopointer_test

import (
	twopointer "leetcode/twoPointer"
	"testing"
)

var testWord = []struct {
	input    string
	expected string
}{
	{"God", "doG"},
	{"God Ding", "doG gniD"},
}

func TestReverseWord(t *testing.T) {
	t.Run("Expected right answer", func(t *testing.T) {
		for _, test := range testWord {

			outArr := twopointer.ReverseWord(test.input)

			if len(test.expected) != len(outArr) {
				t.Error("Test Failed. Expected: ", test.expected, " Output: ", outArr)
			}

			if test.expected != outArr {
				t.Error("Test Failed. Expected:", test.expected, " Output:", outArr)
			}

		}
	})
}
