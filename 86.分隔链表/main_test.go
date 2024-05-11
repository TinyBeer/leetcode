package main

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "head = [1,4,3,2,5,2], x = 3",
			args: args{
				head: NewListFromSlice([]int{1, 4, 3, 2, 5, 2}),
				x:    3,
			},
			want: NewListFromSlice([]int{1, 2, 2, 4, 3, 5}),
		},
		{
			name: "head = [2,1], x = 2",
			args: args{
				head: NewListFromSlice([]int{2, 1}),
				x:    2,
			},
			want: NewListFromSlice([]int{1, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.args.head, tt.args.x)
			if !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("partition() = %v, want %v", got.String(), tt.want.String())
			}
		})
	}
}
