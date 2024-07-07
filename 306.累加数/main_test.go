package main

import "testing"

func Test_isAdditiveNumber(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "112358",
			args: args{
				num: "112358",
			},
			want: true,
		},
		{
			name: "199100199",
			args: args{
				num: "199100199",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAdditiveNumber(tt.args.num); got != tt.want {
				t.Errorf("isAdditiveNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
