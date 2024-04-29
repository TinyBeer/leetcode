package main

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "lists = [[1,4,5],[1,3,4],[2,6]]",
			args: args{
				lists: []*ListNode{NewList(1, 4, 5), NewList(1, 3, 4), NewList(2, 6)},
			},
			want: NewList(1, 1, 2, 3, 4, 4, 5, 6),
		},
		{
			name: "lists = []",
			args: args{
				lists: []*ListNode{},
			},
			want: NewList(),
		},
		{
			name: "lists = [[]]",
			args: args{
				lists: []*ListNode{NewList()},
			},
			want: NewList(),
		},
		{
			name: "lists = [[1,2,3],[4,5,6,7]]",
			args: args{
				lists: []*ListNode{NewList(1, 2, 3), NewList(4, 5, 6, 7)},
			},
			want: NewList(1, 2, 3, 4, 5, 6, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeKLists(tt.args.lists).Show()
			want := tt.want.Show()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("mergeKLists() = %v, want %v", got, want)
			}
		})
	}
}
