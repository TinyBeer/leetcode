package main

import (
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "head = [1,2,3,4]",
			args: args{
				head: NewList(1, 2, 3, 4),
			},
			want: NewList(2, 1, 4, 3),
		},
		{
			name: "head = []",
			args: args{
				head: NewList(),
			},
			want: NewList(),
		},
		{
			name: "head = [1]",
			args: args{
				head: NewList(1),
			},
			want: NewList(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := swapPairs(tt.args.head).Show()
			want := tt.want.Show()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("mergeKLists() = %v, want %v", got, want)
			}
		})
	}
}
