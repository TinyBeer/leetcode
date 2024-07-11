package main

import "testing"

func Test_nthSuperUglyNumber(t *testing.T) {
	type args struct {
		n      int
		primes []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 12, primes = [2,7,13,19]",
			args: args{
				n:      12,
				primes: []int{2, 7, 13, 19},
			},
			want: 32,
		},
		{
			name: "n = 1, primes = [2,3,5]",
			args: args{
				n:      1,
				primes: []int{2, 3, 5},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthSuperUglyNumber(tt.args.n, tt.args.primes); got != tt.want {
				t.Errorf("nthSuperUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
