package main

import "testing"

func Test_titleToNumber(t *testing.T) {
	type args struct {
		columnTitle string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `columnTitle = "A"`,
			args: args{
				columnTitle: "A",
			},
			want: 1,
		},
		{
			name: `columnTitle = "AB"`,
			args: args{
				columnTitle: "AB",
			},
			want: 28,
		},
		{
			name: `columnTitle = "ZY"`,
			args: args{
				columnTitle: "ZY",
			},
			want: 701,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.columnTitle); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
