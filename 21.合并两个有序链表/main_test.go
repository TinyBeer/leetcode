package main

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "l1 = [1,2,4], l2 = [1,3,4]",
			args: args{
				list1: NewList([]int{1, 2, 4}),
				list2: NewList([]int{1, 3, 4}),
			},
			want: NewList([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name: "l1 = [], l2 = [0]",
			args: args{
				list1: NewList([]int{}),
				list2: NewList([]int{0}),
			},
			want: NewList([]int{0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.list1, tt.args.list2).Show()
			want := tt.want.Show()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, want)
			}
		})
	}
}
