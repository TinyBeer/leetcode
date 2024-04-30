package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: `s = "barfoothefoobarman", words = ["foo","bar"]`,
			args: args{
				s:     "barfoothefoobarman",
				words: []string{"foo", "bar"},
			},
			want: []int{0, 9},
		},
		{
			name: `s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]`,
			args: args{
				s:     "wordgoodgoodgoodbestword",
				words: []string{"word", "good", "best", "word"},
			},
			want: nil,
		},
		{
			name: `s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]`,
			args: args{
				s:     "barfoofoobarthefoobarman",
				words: []string{"bar", "foo", "the"},
			},
			want: []int{6, 9, 12},
		},
		{
			name: `s = "a", words = ["a"]`,
			args: args{
				s:     "a",
				words: []string{"a"},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSubstring(tt.args.s, tt.args.words)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
