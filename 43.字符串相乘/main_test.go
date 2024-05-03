package main

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `num1 = "2", num2 = "3"`,
			args: args{
				num1: "2",
				num2: "3",
			},
			want: "6",
		},
		{
			name: `num1 = "123", num2 = "456"`,
			args: args{
				num1: "123",
				num2: "456",
			},
			want: "56088",
		},
		{
			name: `num1 = "498828660196", num2 = "840477629533"`,
			args: args{
				num1: "498828660196",
				num2: "840477629533",
			},
			want: "419254329864656431168468",
		},
		{
			name: `num1 = "925101087184894", num2 = "3896737933784656127"`,
			args: args{
				num1: "925101087184894",
				num2: "3896737933784656127",
			},
			want: "3604876499018802870538090258945538",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
