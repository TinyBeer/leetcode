package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "x = 121",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "x = -121",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			name: "x = 10",
			args: args{
				x: 10,
			},
			want: false,
		},
		{
			name: "x = 11",
			args: args{
				x: 11,
			},
			want: true,
		},
		{
			name: "x = -1",
			args: args{
				x: -1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
