package main

import (
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{

		{
			name: "inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]",
			args: args{
				inorder:   []int{9, 3, 15, 20, 7},
				postorder: []int{9, 15, 7, 20, 3},
			},
			want: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			name: "preorder = [-1], inorder = [-1]",
			args: args{
				inorder:   []int{-1},
				postorder: []int{-1},
			},
			want: &TreeNode{Val: -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PreOderTraversal(buildTree(tt.args.inorder, tt.args.postorder))
			want := PreOderTraversal(tt.want)
			if got != want {
				t.Errorf("buildTree() = %v, want %v", got, want)
			}
		})
	}
}
