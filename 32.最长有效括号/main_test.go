package main

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `s = "(()"`,
			args: args{
				s: "(()",
			},
			want: 2,
		},
		{
			name: `s = ")()())"`,
			args: args{
				s: ")()())",
			},
			want: 4,
		},
		{
			name: `s = ""`,
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: `s = "()(()"`,
			args: args{
				s: "()(()",
			},
			want: 2,
		},
		{
			name: `s = ""())""`,
			args: args{
				s: "())",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
