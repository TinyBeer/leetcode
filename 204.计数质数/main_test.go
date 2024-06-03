package main

import "testing"

func Test_countPrimes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 10",
			args: args{
				n: 10,
			},
			want: 4,
		},
		{
			name: "n = 0",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "n = 1",
			args: args{
				n: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimes(tt.args.n); got != tt.want {
				t.Errorf("countPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
