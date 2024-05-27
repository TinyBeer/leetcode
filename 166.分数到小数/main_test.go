package main

import "testing"

func Test_fractionToDecimal(t *testing.T) {
	type args struct {
		numerator   int
		denominator int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "numerator = 1, denominator = 2",
			args: args{
				numerator:   1,
				denominator: 2,
			},
			want: "0.5",
		},
		{
			name: "numerator = 2, denominator = 1",
			args: args{
				numerator:   2,
				denominator: 1,
			},
			want: "2",
		},
		{
			name: "numerator = 4, denominator = 333",
			args: args{
				numerator:   4,
				denominator: 333,
			},
			want: "0.(012)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fractionToDecimal(tt.args.numerator, tt.args.denominator); got != tt.want {
				t.Errorf("fractionToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
