package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_permuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "nums = [1,1,2]",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1}},
		},
		{
			name: "nums = [1,2,3]",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permuteUnique(tt.args.nums)
			sort.Slice(got, func(i, j int) bool {
				for x := 0; x < len(got[i]); x++ {
					if got[i][x] == got[j][x] {
						continue
					}
					return got[i][x] < got[j][x]
				}
				return false
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
