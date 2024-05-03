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
			name: `s = "aa", p = "*"`,
			args: args{
				s: "aa",
				p: "*",
			},
			want: true,
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
			name: `s = "cb", p = "?a"`,
			args: args{
				s: "cb",
				p: "?a",
			},
			want: false,
		},
		{
			name: `s = "bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab", p = "b*b*ab**ba*b**b***bba"`,
			args: args{
				s: "bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab",
				p: "b*b*ab**ba*b**b***bba",
			},
			want: false,
		},
		{
			name: `s = "adceb", p = "*a*b"`,
			args: args{
				s: "adceb",
				p: "*a*b",
			},
			want: true,
		},
		{
			name: `s = " ", p = "******"`,
			args: args{
				s: "",
				p: "******",
			},
			want: true,
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
