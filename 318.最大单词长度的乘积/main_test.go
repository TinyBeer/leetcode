package main

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `words = ["abcw","baz","foo","bar","xtfn","abcdef"]`,
			args: args{words: []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}},
			want: 16,
		},
		{
			name: `words = ["a","ab","abc","d","cd","bcd","abcd"]`,
			args: args{words: []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}},
			want: 4,
		},
		{
			name: `words = ["a","aa","aaa","aaaa"]`,
			args: args{words: []string{"a", "aa", "aaa", "aaaa"}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.words); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
