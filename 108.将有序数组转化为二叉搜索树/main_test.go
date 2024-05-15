package main

import (
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: "nums = [-10,-3,0,5,9]",
			args: args{
				nums: []int{-10, -3, 0, 5, 9},
			},
			want: []*TreeNode{
				{
					Val: 0,
					Left: &TreeNode{
						Val:  -3,
						Left: &TreeNode{Val: -10},
					},
					Right: &TreeNode{
						Val:  9,
						Left: &TreeNode{Val: 5},
					},
				},
				{
					Val: 0,
					Left: &TreeNode{
						Val:   -10,
						Right: &TreeNode{Val: -3},
					},
					Right: &TreeNode{
						Val:   5,
						Right: &TreeNode{Val: 9},
					},
				},
			},
		},
		{
			name: "nums = [1,3]",
			args: args{
				nums: []int{1, 3},
			},
			want: []*TreeNode{
				{
					Val:   1,
					Right: &TreeNode{Val: 3},
				},
				{
					Val:  3,
					Left: &TreeNode{Val: 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PreOderTraversal(sortedArrayToBST(tt.args.nums))
			for _, w := range tt.want {
				want := PreOderTraversal(w)
				if want == got {
					return
				}
			}
			t.Errorf("sortedArrayToBST() = %v, want", got)
		})
	}
}
