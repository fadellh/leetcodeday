package array

import "testing"

func Test_isValidSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				board: [][]byte{
					{53, 51, 46, 46, 55, 46, 46, 46, 46},
					{54, 46, 46, 49, 57, 53, 46, 46, 46},
					{46, 57, 56, 46, 46, 46, 46, 54, 46},
					{56, 46, 46, 46, 54, 46, 46, 46, 51},
					{52, 46, 46, 56, 46, 51, 46, 46, 49},
					{55, 46, 46, 46, 50, 46, 46, 46, 54},
					{46, 54, 46, 46, 46, 46, 50, 56, 46},
					{46, 46, 46, 52, 49, 57, 46, 46, 53},
					{46, 46, 46, 46, 56, 46, 46, 55, 57},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSudoku(tt.args.board); got != tt.want {
				t.Errorf("isValidSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}
