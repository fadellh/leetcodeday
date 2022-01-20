package twopointer_test

import (
	twopointer "leetcode/leetTopic/twoPointer"
	"testing"
)

var testString = []struct {
	input    []byte
	expected []byte
}{
	{[]byte{123, 221}, []byte{221, 123}},
}

func TestReverseString(t *testing.T) {
	t.Run("Expected right answer", func(t *testing.T) {
		for _, test := range testString {

			outArr := twopointer.ReverseString(test.input)

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
