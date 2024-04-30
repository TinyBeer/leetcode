package main

import (
	"sort"
	"testing"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	type exps struct {
		k    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		exps exps
	}{
		{
			name: "nums = [3,2,2,3], val = 3",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			exps: exps{
				k:    2,
				nums: []int{2, 2},
			},
		},
		{
			name: "nums = [0,1,2,2,3,0,4,2], val = 2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			exps: exps{
				k:    5,
				nums: []int{0, 1, 3, 0, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.args.nums, tt.args.val)
			if got != tt.exps.k {
				t.Errorf("removeElement() = %v, want %v", got, tt.exps.k)
			}
			sort.Ints(tt.args.nums[:got])
			sort.Ints(tt.exps.nums)
			for i := 0; i < tt.exps.k; i++ {
				if tt.args.nums[i] != tt.exps.nums[i] {
					t.Errorf("nums = %v, want %v", tt.exps.nums, tt.exps.nums)
				}
			}
		})
	}
}
