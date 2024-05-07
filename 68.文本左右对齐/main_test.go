package main

import (
	"reflect"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: `words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16`,
			args: args{
				words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
				maxWidth: 16,
			},
			want: []string{
				"This    is    an",
				"example  of text",
				"justification.  "},
		},
		{
			name: `words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16`,
			args: args{
				words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
				maxWidth: 16,
			},
			want: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify() = %q, want %q", got, tt.want)
			}
		})
	}
}
