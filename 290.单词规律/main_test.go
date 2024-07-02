package main

import "testing"

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		s       string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `pattern = "abba", s = "dog cat cat dog"`,
			args: args{
				pattern: "abba",
				s:       "dog cat cat dog",
			},
			want: true,
		},
		{
			name: `pattern = "abba", s = "dog cat cat fish"`,
			args: args{
				pattern: "abba",
				s:       "dog cat cat fish",
			},
			want: false,
		},
		{
			name: `pattern = "aaaa", s = "dog cat cat dog"`,
			args: args{
				pattern: "aaaa",
				s:       "dog cat cat fish",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.s); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
