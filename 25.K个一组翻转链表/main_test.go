package main

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "head = [1,2,3,4,5], k = 2",
			args: args{
				head: NewList(1, 2, 3, 4, 5),
				k:    2,
			},
			want: NewList(2, 1, 4, 3, 5),
		},
		{
			name: "head = [1,2,3,4,5], k = 2",
			args: args{
				head: NewList(1, 2, 3, 4, 5),
				k:    3,
			},
			want: NewList(3, 2, 1, 4, 5),
		},
		{
			name: "head = [1,2], k = 2",
			args: args{
				head: NewList(1, 2),
				k:    2,
			},
			want: NewList(2, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseKGroup(tt.args.head, tt.args.k).Show()
			want := tt.want.Show()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("mergeKLists() = %v, want %v", got, want)
			}
		})
	}
}
