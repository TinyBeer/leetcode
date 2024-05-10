package main

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "head = [1,2,3,3,4,4,5]",
			args: args{
				head: NewListFromSlice([]int{1, 2, 3, 3, 4, 4, 5}),
			},
			want: NewListFromSlice([]int{1, 2, 5}),
		},
		{
			name: "head = [1,1,1,2,3]",
			args: args{
				head: NewListFromSlice([]int{1, 1, 1, 2, 3}),
			},
			want: NewListFromSlice([]int{2, 3}),
		},
		{
			name: "head = [1,1]",
			args: args{
				head: NewListFromSlice([]int{1, 1}),
			},
			want: NewListFromSlice(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("deleteDuplicates() = %v, want %v", got.String(), tt.want.String())
			}
		})
	}
}
