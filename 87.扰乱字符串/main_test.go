package main

import "testing"

func Test_isScramble(t *testing.T) {
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
			name: `s1 = "great", s2 = "rgeat"`,
			args: args{
				s1: "great",
				s2: "rgeat",
			},
			want: true,
		},
		{
			name: `s1 = "abcde", s2 = "caebd"`,
			args: args{
				s1: "abcde",
				s2: "caebd",
			},
			want: false,
		},
		{
			name: `s1 = "a", s2 = "a"`,
			args: args{
				s1: "a",
				s2: "a",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isScramble(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isScramble() = %v, want %v", got, tt.want)
			}
		})
	}
}
