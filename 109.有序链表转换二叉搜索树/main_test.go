package main

import (
	"testing"
)

func Test_sortedListToBST(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: "head = [-10,-3,0,5,9]",
			args: args{
				head: GenNodeList(-10, -3, 0, 5, 9),
			},
			want: []*TreeNode{{
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
			name: "head = []",
			args: args{
				head: nil,
			},
			want: []*TreeNode{nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := PreOderTraversal(sortedListToBST(tt.args.head))
			for _, w := range tt.want {
				if got == PreOderTraversal(w) {
					return
				}
			}
			t.Errorf("sortedListToBST() = %v", got)

		})
	}
}
