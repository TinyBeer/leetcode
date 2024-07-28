package main

import "testing"

func Test_isPowerOfFour(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "n = 16",
			args: args{
				n: 16,
			},
			want: true,
		},
		{
			name: "n = 5",
			args: args{
				n: 5,
			},
			want: false,
		},
		{
			name: "n = 1",
			args: args{
				n: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}
