package slidingWindow

import "testing"

func TestCheckInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ab",
			args: args{
				s1: "ab",
				s2: "eidbaoab",
			},
			want: true,
		},
		{
			name: "ab2",
			args: args{
				s1: "ab",
				s2: "eidboaoo",
			},
			want: false,
		},
		{
			name: "adc",
			args: args{
				s1: "adc",
				s2: "dcda",
			},
			want: true,
		},
		{
			name: "hello",
			args: args{
				s1: "hello",
				s2: "ooolleeeh",
			},
			want: false,
		},
		{
			name: "hello2",
			args: args{
				s1: "hello",
				s2: "ooolleoooleh",
			},
			want: false,
		},
		{
			name: "abc",
			args: args{
				s1: "abc",
				s2: "bbbca",
			},
			want: true,
		},
		{
			name: "abc2",
			args: args{
				s1: "abc",
				s2: "ccccbbbbaaaa",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckInclusionLeet(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CheckInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
