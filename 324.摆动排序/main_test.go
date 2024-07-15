package main

import (
	"reflect"
	"testing"
)

func Test_wiggleSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [1,5,1,1,6,4]",
			args: args{
				nums: []int{1, 5, 1, 1, 6, 4},
			},
			want: []int{1, 6, 1, 5, 1, 4},
		},
		{
			name: "nums = [1,3,2,2,3,1]",
			args: args{
				nums: []int{1, 3, 2, 2, 3, 1},
			},
			want: []int{2, 3, 1, 3, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wiggleSort(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("wiggleSort() %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
