package binarysearch

import "testing"

func TestSearchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "succesmid_equal_target",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			want: 2,
		},
		{
			name: "succes mid > target",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 6,
			},
			want: 3,
		},
		{
			name: "succes mid < target",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 3,
			},
			want: 1,
		},
		{
			name: "success no value, return index would it be",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("SearchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
