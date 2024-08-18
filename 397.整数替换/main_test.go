package main

import "testing"

func Test_integerReplacement(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 8",
			args: args{
				n: 8,
			},
			want: 3,
		},
		{
			name: "n = 7",
			args: args{
				n: 7,
			},
			want: 4,
		},
		{
			name: "n = 4",
			args: args{
				n: 4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integerReplacement(tt.args.n); got != tt.want {
				t.Errorf("integerReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
