package main

import "testing"

func Test_bulbSwitch(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 30",
			args: args{
				n: 3,
			},
			want: 1,
		},
		{
			name: "n = 0",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "n = 1",
			args: args{
				n: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bulbSwitch(tt.args.n); got != tt.want {
				t.Errorf("bulbSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}
