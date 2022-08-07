package twopointer

import "testing"

func TestRotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				nums: []int{},
				k:    0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rotate(tt.args.nums, tt.args.k)
		})
	}
}
