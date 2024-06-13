package main

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `s = "3+2*2"`,
			args: args{
				s: "3+2*2",
			},
			want: 7,
		},
		{
			name: `s = "3/2"`,
			args: args{
				s: "3/2",
			},
			want: 1,
		},
		{
			name: `s = " 3+5 / 2 "`,
			args: args{
				s: " 3+5 / 2",
			},
			want: 5,
		},
		{
			name: `s = "1*2-3/4+5*6-7*8+9/10"`,
			args: args{
				s: "1*2-3/4+5*6-7*8+9/10",
			},
			want: -24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
