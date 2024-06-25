package main

import "testing"

func Test_isUgly(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "n = 6",
			args: args{
				n: 6,
			},
			want: true,
		},
		{
			name: "n = 1",
			args: args{
				n: 1,
			},
			want: true,
		},
		{
			name: "n = 14",
			args: args{
				n: 14,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUgly(tt.args.n); got != tt.want {
				t.Errorf("isUgly() = %v, want %v", got, tt.want)
			}
		})
	}
}
