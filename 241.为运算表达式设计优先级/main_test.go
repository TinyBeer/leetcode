package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_diffWaysToCompute(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: `expression = "2-1-1"`,
			args: args{
				expression: "2-1-1",
			},
			want: []int{0, 2},
		},
		{
			name: `expression = "2*3-4*5"`,
			args: args{
				expression: "2*3-4*5",
			},
			want: []int{-34, -14, -10, -10, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diffWaysToCompute(tt.args.expression)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diffWaysToCompute() = %v, want %v", got, tt.want)
			}
		})
	}
}
