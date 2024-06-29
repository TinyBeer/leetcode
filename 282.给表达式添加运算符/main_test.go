package main

import (
	"reflect"
	"testing"
)

func Test_addOperators(t *testing.T) {
	type args struct {
		num    string
		target int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: `num = "123", target = 6`,
			args: args{
				num:    "123",
				target: 6,
			},
			want: []string{"1+2+3", "1*2*3"},
		},
		{
			name: `num = "232", target = 8`,
			args: args{
				num:    "232",
				target: 8,
			},
			want: []string{"2+3*2", "2*3+2"},
		},
		{
			name: `num = "3456237490", target = 9191`,
			args: args{
				num:    "3456237490",
				target: 9191,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addOperators(tt.args.num, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOperators() = %v, want %v", got, tt.want)
			}
		})
	}
}
