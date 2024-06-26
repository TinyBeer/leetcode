package main

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `s = "12"`,
			args: args{
				s: "12",
			},
			want: 2,
		},
		{
			name: `s = "226"`,
			args: args{
				s: "226",
			},
			want: 3,
		},
		{
			name: `s = "06"`,
			args: args{
				s: "06",
			},
			want: 0,
		},
		{
			name: `s = "2101"`,
			args: args{
				s: "2101",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
