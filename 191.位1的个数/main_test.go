package main

import "testing"

func Test_hammingWeight(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 11",
			args: args{
				n: 11,
			},
			want: 3,
		},
		{
			name: "n = 128",
			args: args{
				n: 128,
			},
			want: 1,
		},
		{
			name: "n = 2147483645",
			args: args{
				n: 2147483645,
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.n); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
