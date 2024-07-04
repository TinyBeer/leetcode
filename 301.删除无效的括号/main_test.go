package main

import (
	"reflect"
	"testing"
)

func Test_removeInvalidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: `s = "()())()"`,
			args: args{
				s: "()())()",
			},
			want: []string{"(())()", "()()()"},
		},
		{
			name: `s = "(a)())()"`,
			args: args{
				s: "(a)())()",
			},
			want: []string{"(a())()", "(a)()()"},
		},
		{
			name: `s = ")("`,
			args: args{
				s: ")(",
			},
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeInvalidParentheses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeInvalidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
