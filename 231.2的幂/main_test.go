package main

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "n = 1",
			args: args{
				n: 1,
			},
			want: true,
		},
		{
			name: "n = 16",
			args: args{
				n: 16,
			},
			want: true,
		},
		{
			name: "n = 3",
			args: args{
				n: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
