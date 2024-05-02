package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "candidates = [2,3,6,7], target = 7",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{{2, 2, 3}, {7}},
		},
		{
			name: "candidates = [2,3,5], target = 8",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			want: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name: "candidates = [2], target = 1",
			args: args{
				candidates: []int{2},
				target:     1,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.args.candidates, tt.args.target)
			sort.Slice(got, func(i, j int) bool {
				n := len(got[i])
				m := len(got[j])
				for k := 0; k < n && k < m; k++ {
					if got[i][k] != got[j][k] {
						return got[i][k] < got[j][k]
					}
				}
				return n < m
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
