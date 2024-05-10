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
			name: "head = [1,1,2]",
			args: args{
				head: NewListFromSlice([]int{1, 1, 2}),
			},
			want: NewListFromSlice([]int{1, 2}),
		},
		{
			name: "head = [1,1,2,3,3]",
			args: args{
				head: NewListFromSlice([]int{1, 1, 2, 3, 3}),
			},
			want: NewListFromSlice([]int{1, 2, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
