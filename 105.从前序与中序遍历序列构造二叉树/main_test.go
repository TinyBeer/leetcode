package main

import (
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
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
				preorder: []int{-1},
				inorder:  []int{-1},
			},
			want: &TreeNode{Val: -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PreOderTraversal(buildTree(tt.args.preorder, tt.args.inorder))
			want := PreOderTraversal(tt.want)
			if got != want {
				t.Errorf("buildTree() = %v, want %v", got, want)
			}
		})
	}
}
