package main

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	type expect struct {
		k    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		exp  expect
	}{
		{
			name: "nums = [1,1,2]",
			args: args{nums: []int{1, 1, 2}},
			exp: expect{
				k:    2,
				nums: []int{1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			if got != tt.exp.k {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.exp.k)
			}
			for i := 0; i < tt.exp.k; i++ {
				if tt.args.nums[i] != tt.exp.nums[i] {
					t.Errorf("nums = %v, want %v", tt.exp.nums, tt.exp.nums)
				}
			}
		})
	}
}
