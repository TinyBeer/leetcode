package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_wordBreak(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: `s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]`,
			args: args{
				s:        "catsanddog",
				wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			},
			want: []string{"cats and dog", "cat sand dog"},
		},
		{
			name: `s = "pineapplepenapple", wordDict = ["apple","pen","applepen","pine","pineapple"]`,
			args: args{
				s:        "pineapplepenapple",
				wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			},
			want: []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"},
		},
		{
			name: `s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]`,
			args: args{
				s:        "catsandog",
				wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreak(tt.args.s, tt.args.wordDict)
			sort.Slice(got, func(i, j int) bool {
				return got[i] < got[j]
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i] < tt.want[j]
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
