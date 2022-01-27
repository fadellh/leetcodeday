package slidingWindow_test

import (
	sd "leetcode/leetTopic/slidingWindow"
	"testing"
)

func Test_longestSubstring(t *testing.T) {
	type parameters struct {
		s string
	}
	type testCase struct {
		name string
		args parameters
		want int
	}
	tests := []testCase{
		{
			name: "test 0 abcabcab",
			args: parameters{"abcabcab"},
			want: 3,
		},
		{
			name: "test 1 bbbbbb",
			args: parameters{"bbbbbb"},
			want: 1,
		},
		{
			name: "test 2 pwwkew",
			args: parameters{"pwwkew"},
			want: 3,
		}, {
			name: "test 3 bbbcdb",
			args: parameters{"bbbcdb"},
			want: 3,
		},
		{
			name: "test 4 aab",
			args: parameters{"aab"},
			want: 2,
		},
		{
			name: "test 5 dvdf",
			args: parameters{"dvdf"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sd.LongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("LongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
