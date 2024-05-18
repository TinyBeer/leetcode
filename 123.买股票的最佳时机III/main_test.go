package main

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "prices = [3,3,5,0,0,3,1,4]",
			args: args{
				prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			},
			want: 6,
		},
		{
			name: "prices = [1,2,3,4,5]",
			args: args{
				prices: []int{1, 2, 3, 4, 5},
			},
			want: 4,
		},
		{
			name: "prices = [7,6,4,3,1] ",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
		{
			name: "prices = [1]",
			args: args{
				prices: []int{1},
			},
			want: 0,
		},
		{
			name: "prices = [14,9,10,12,4,8,1,16]",
			args: args{
				prices: []int{14, 9, 10, 12, 4, 8, 1, 16},
			},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
