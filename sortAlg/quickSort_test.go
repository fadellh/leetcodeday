package sortalg

import (
	"reflect"
	"testing"
)

func Test_qsort(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				a: []int{2, 1},
			},
			want: []int{1, 2},
		},
		{
			name: "2",
			args: args{
				a: []int{5, 6, 7, 8, 1, 4, 3, 2},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := qsort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("qsort() = %v, want %v", got, tt.want)
			}
		})
	}
}
