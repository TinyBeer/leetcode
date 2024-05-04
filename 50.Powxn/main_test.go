package main

import (
	"fmt"
	"testing"
)

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "x = 2.00000, n = 10",
			args: args{
				x: 2,
				n: 10,
			},
			want: 1024,
		},
		{
			name: "x = 2.10000, n = 3",
			args: args{
				x: 2.1,
				n: 3,
			},
			want: 9.26100,
		},
		{
			name: "x = 2.00000, n = -2",
			args: args{
				x: 2,
				n: -2,
			},
			want: 0.25000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); fmt.Sprintf("%.5f", got) != fmt.Sprintf("%.5f", tt.want) {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
