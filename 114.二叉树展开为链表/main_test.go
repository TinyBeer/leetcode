package main

import "testing"

func Test_flatten(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "root = [1,2,5,3,4,null,6]",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 4},
					},
					Right: &TreeNode{
						Val:   5,
						Right: &TreeNode{Val: 6},
					},
				},
			},
			want: "[1,null,2,null,3,null,4,null,5,null,6,null,null]",
		},
		{
			name: "root = []",
			args: args{
				root: nil,
			},
			want: "[null]",
		},
		{
			name: "root = [0]",
			args: args{
				root: &TreeNode{Val: 0},
			},
			want: "[0,null,null]",
		},
		{
			name: "root = [1,2,null,3]",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 3},
					},
				},
			},
			want: "[1,null,2,null,3,null,null]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.args.root)
			got := PreOrder(tt.args.root)
			if got != tt.want {
				t.Errorf("flatten()  %v, want %v", got, tt.want)
			}
		})
	}
}
