package main

import "testing"

func Test_increasingTriplet(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nums = [1,2,3,4,5]",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: true,
		},
		{
			name: "nums = [5,4,3,2,1]",
			args: args{
				nums: []int{5, 4, 3, 2, 1},
			},
			want: false,
		},
		{
			name: "nums = [2,1,5,0,4,6]",
			args: args{
				nums: []int{2, 1, 5, 0, 4, 6},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingTriplet(tt.args.nums); got != tt.want {
				t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}
