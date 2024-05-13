package main

import (
	"reflect"
	"testing"
)

func Test_generateTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: "n = 3",
			args: args{
				n: 3,
			},
			want: []*TreeNode{
				GenerateTree(1, -1, 2, -1, 3),
				GenerateTree(1, -1, 3, 2),
				GenerateTree(2, 1, 3),
				GenerateTree(3, 1, -1, -1, 2),
				GenerateTree(3, 2, -1, 1),
			},
		},
		{
			name: "n=1",
			args: args{
				n: 1,
			},
			want: generateTrees(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateTrees(tt.args.n); !reflect.DeepEqual(PreOderTraversal(got), PreOderTraversal(tt.want)) {
				t.Errorf("generateTrees()\ngot = %v\nwant %v", PreOderTraversal(got), PreOderTraversal(tt.want))
			}
		})
	}
}
