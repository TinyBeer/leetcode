package main

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		k      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "k = 2, prices = [2,4,1]",
			args: args{
				k:      2,
				prices: []int{2, 4, 1},
			},
			want: 2,
		},
		{
			name: "k = 2, prices = [3,2,6,5,0,3]",
			args: args{
				k:      2,
				prices: []int{3, 2, 6, 5, 0, 3},
			},
			want: 7,
		},
		{
			name: "k = 1, prices = [3,2,6,5,0,3]",
			args: args{
				k:      1,
				prices: []int{3, 2, 6, 5, 0, 3},
			},
			want: 4,
		},
		{
			name: "k = 3, prices = [3,2,6,5,0,3,4,8,2]",
			args: args{
				k:      3,
				prices: []int{3, 2, 6, 5, 0, 3, 4, 8, 2, 5, 5},
			},
			want: 15,
		},
		{
			name: "k = 2, prices = [1,2,4,2,5,7,2,4,9,0]",
			args: args{
				k:      2,
				prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
