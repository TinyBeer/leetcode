package main

import "testing"

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "n = 00000010100101000001111010011100",
			args: args{
				num: 0b00000010100101000001111010011100,
			},
			want: 964176192,
		},
		{
			name: "n = 11111111111111111111111111111101",
			args: args{
				num: 0b11111111111111111111111111111101,
			},
			want: 3221225471,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.num); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
