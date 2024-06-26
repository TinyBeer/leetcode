package main

import "testing"

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [2,2,3,2]",
			args: args{
				nums: []int{2, 2, 3, 2},
			},
			want: 3,
		},
		{
			name: "nums = [0,1,0,1,0,1,99]",
			args: args{
				nums: []int{0, 1, 0, 1, 0, 1, 99},
			},
			want: 99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
