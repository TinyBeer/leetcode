package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "n = 4, k = 2",
			args: args{
				n: 4,
				k: 2,
			},
			want: [][]int{
				{2, 4},
				{3, 4},
				{2, 3},
				{1, 2},
				{1, 3},
				{1, 4},
			},
		},
		{
			name: "n = 1, k = 1",
			args: args{
				n: 1,
				k: 1,
			},
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combine(tt.args.n, tt.args.k)
			sort.Slice(tt.want, func(i, j int) bool {
				for l := 0; l < len(tt.want[i]); l++ {
					if tt.want[i][l] == tt.want[j][l] {
						continue
					}
					return tt.want[i][l] < tt.want[j][l]
				}
				return false
			})
			sort.Slice(got, func(i, j int) bool {
				for l := 0; l < len(got[i]); l++ {
					if got[i][l] == got[j][l] {
						continue
					}
					return got[i][l] < got[j][l]
				}
				return false
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}
