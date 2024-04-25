package main

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "babad",
			args: args{
				s: "babad",
			},
			want: "bab",
		},

		{
			name: "cbbd",
			args: args{
				s: "bb",
			},
			want: "bb",
		},
		{
			name: "aacabdkacaa",
			args: args{
				s: "aacabdkacaa",
			},
			want: "aca",
		},
		{
			name: "aaaaa",
			args: args{
				s: "aaaaa",
			},
			want: "aaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
