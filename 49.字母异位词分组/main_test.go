package main

import (
	"encoding/json"
	"reflect"
	"sort"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: `strs = ["eat", "tea", "tan", "ate", "nat", "bat"]`,
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{{"bat"}, {"tan", "nat"}, {"eat", "tea", "ate"}},
		},
		{
			name: `strs = [""]`,
			args: args{
				strs: []string{""},
			},
			want: [][]string{{""}},
		},
		{
			name: `strs = ["a"]`,
			args: args{
				strs: []string{"a"},
			},
			want: [][]string{{"a"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.args.strs)
			sort.Slice(got, func(i, j int) bool {
				return len(got[i]) < len(got[j])
			})
			gjs, _ := json.Marshal(got)
			wjs, _ := json.Marshal(tt.want)
			if !reflect.DeepEqual(gjs, wjs) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
