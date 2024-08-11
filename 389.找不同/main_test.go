package main

import "testing"

func Test_findTheDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: `s = "abcd", t = "abcde"`,
			args: args{
				s: "abcd",
				t: "abcde",
			},
			want: 'e',
		},
		{
			name: `s = "", t = "y"`,
			args: args{
				s: "",
				t: "y",
			},
			want: 'y',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifference(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
