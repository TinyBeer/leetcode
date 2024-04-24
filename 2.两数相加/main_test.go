package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func generateList(nums []int) *ListNode {
	head := &ListNode{}
	helper := head
	for _, v := range nums {
		helper.Next = &ListNode{Val: v}
		helper = helper.Next
	}
	return head.Next
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				l1: generateList([]int{2, 4, 3}),
				l2: generateList([]int{5, 6, 4}),
			},
			want: generateList([]int{7, 0, 8}),
		},

		{
			name: "",
			args: args{
				l1: generateList([]int{0}),
				l2: generateList([]int{0}),
			},
			want: generateList([]int{0}),
		},
		{
			name: "",
			args: args{
				l1: generateList([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: generateList([]int{9, 9, 9, 9}),
			},
			want: generateList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.args.l1, tt.args.l2)
			gjs, _ := json.Marshal(got)
			wjs, _ := json.Marshal(tt.want)
			if !reflect.DeepEqual(gjs, wjs) {
				t.Errorf("addTwoNumbers() = %v, want %v", string(gjs), string(wjs))
			}
		})
	}
}
