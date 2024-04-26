package main

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `s = "42"`,
			args: args{
				s: "42",
			},
			want: 42,
		},
		{
			name: `s = "   -42"`,
			args: args{
				s: "   -42",
			},
			want: -42,
		},
		{
			name: `s = "4193 with words"`,
			args: args{
				s: "4193 with words",
			},
			want: 4193,
		},
		{
			name: `s = "words and 987"`,
			args: args{
				s: "words and 987",
			},
			want: 0,
		},
		{
			name: `s = "-91283472332"`,
			args: args{
				s: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: `s = "+-12"`,
			args: args{
				s: "+-12",
			},
			want: 0,
		},
		{
			name: `s = "9223372036854775808"`,
			args: args{
				s: "9223372036854775808",
			},
			want: 2147483647,
		},
		{
			name: `s = "-2147483648"`,
			args: args{
				s: "-2147483648",
			},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
