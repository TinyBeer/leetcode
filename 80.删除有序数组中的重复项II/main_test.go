package main

import (
	"fmt"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		final []int
	}{
		{
			name:  "nums = [1,1,1,2,2,3]",
			args:  args{nums: []int{1, 1, 1, 2, 2, 3}},
			want:  5,
			final: []int{1, 1, 2, 2, 3},
		},
		{
			name:  "nums = [0,0,1,1,1,1,2,3,3]",
			args:  args{nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3}},
			want:  7,
			final: []int{0, 0, 1, 1, 2, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			fmt.Println(tt.args.nums)
			// if !reflect.DeepEqual(tt.args.nums, tt.final) {
			// 	t.Errorf("removeDuplicates()  %v, want %v", tt.args.nums, tt.final)
			// }
		})
	}
}
