package main

import (
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "head = [1,2,3,4,5], left = 2, right = 4",
			args: args{
				head:  NewListFromSlice([]int{1, 2, 3, 4, 5}),
				left:  2,
				right: 4,
			},
			want: NewListFromSlice([]int{1, 4, 3, 2, 5}),
		},
		{
			name: "head = [5], left = 1, right = 1",
			args: args{
				head:  NewListFromSlice([]int{1}),
				left:  1,
				right: 1,
			},
			want: NewListFromSlice([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("reverseBetween() = %v, want %v", got.String(), tt.want.String())
			}
		})
	}
}
