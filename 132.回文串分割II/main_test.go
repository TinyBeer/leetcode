package main

import "testing"

func Test_minCut(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `s = "aab"`,
			args: args{
				s: "aab",
			},
			want: 1,
		},
		{
			name: "a",
			args: args{
				s: "a",
			},
			want: 0,
		},
		{
			name: "ab",
			args: args{
				s: "ab",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCut(tt.args.s); got != tt.want {
				t.Errorf("minCut() = %v, want %v", got, tt.want)
			}
		})
	}
}
