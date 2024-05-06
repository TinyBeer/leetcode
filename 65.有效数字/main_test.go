package main

import "testing"

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `s = "0"`,
			args: args{
				s: "0",
			},
			want: true,
		},
		{
			name: `s = "e"`,
			args: args{
				s: "e",
			},
			want: false,
		},
		{
			name: `s = "."`,
			args: args{
				s: ".",
			},
			want: false,
		},
		{
			name: `s = "6e6.5"`,
			args: args{
				s: "6e6.5",
			},
			want: false,
		},
		{
			name: `s = "0e1"`,
			args: args{
				s: "0e1",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
