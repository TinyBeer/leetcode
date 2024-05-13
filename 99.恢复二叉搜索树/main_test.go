package main

import (
	"testing"
)

func Test_recoverTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "root = [1,3,null,null,2]",
			args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}}}},
			want: "[3,1,null,null,2]",
		},
		{
			name: "root = [3,1,4,null,null,2]",
			args: args{root: &TreeNode{Val: 3, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 2}}}},
			want: "[2,1,4,null,null,3]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTree(tt.args.root)
			// fmt.Println(PreOderTraversal([]*TreeNode{tt.args.root}))
		})
	}
}
