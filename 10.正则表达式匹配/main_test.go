package main

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `s = "aa", p = "a"`,
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: `s = "aa", p = "a*"`,
			args: args{
				s: "aa",
				p: "a*",
			},
			want: true,
		},
		{
			name: `s = "ab", p = ".*"`,
			args: args{
				s: "ab",
				p: ".*",
			},
			want: true,
		},
		{
			name: `s = "aaa", p = "a*a"`,
			args: args{
				s: "aaa",
				p: "a*a",
			},
			want: true,
		},
		{
			name: `s = "aaa", p = "a.a"`,
			args: args{
				s: "aaa",
				p: "a.a",
			},
			want: true,
		},
		{
			name: `s = "aac", p = ".*c"`,
			args: args{
				s: "aac",
				p: ".*c",
			},
			want: true,
		},
		{
			name: `s = "aac", p = ".*b"`,
			args: args{
				s: "aac",
				p: ".*b",
			},
			want: false,
		},
		{
			name: `s = "aab", p = "c*a*b"`,
			args: args{
				s: "aab",
				p: "c*a*b",
			},
			want: true,
		},
		{
			name: `s = "a", p = ".*..a*"`,
			args: args{
				s: "a",
				p: ".*..a*",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
