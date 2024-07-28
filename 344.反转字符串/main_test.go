package main

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: `s = "hello"`,
			args: args{
				s: []byte("hello"),
			},
			want: []byte("olleh"),
		},
		{
			name: `s = "hannah"`,
			args: args{
				s: []byte("hannah"),
			},
			want: []byte("hannah"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("reverseString()  %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}
