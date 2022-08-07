package array

import "testing"

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				ransomNote: "bbbbgg",
				magazine:   "bbbbgg",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				ransomNote: "bg",
				magazine:   "bbbbgg",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				ransomNote: "a",
				magazine:   "b",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
