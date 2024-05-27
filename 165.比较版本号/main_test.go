package main

import "testing"

func Test_compareVersion(t *testing.T) {
	type args struct {
		version1 string
		version2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `version1 = "1.2", version2 = "1.10"`,
			args: args{
				version1: "1.2",
				version2: "1.10",
			},
			want: -1,
		},
		{
			name: `version1 = "1.01", version2 = "1.001"`,
			args: args{
				version1: "1.01",
				version2: "1.001",
			},
			want: 0,
		},
		{
			name: `version1 = "1.0", version2 = "1.0.0.0"`,
			args: args{
				version1: "1.0",
				version2: "1.0.0.0",
			},
			want: 0,
		},
		{
			name: `version1 = "123", version2 = "321"`,
			args: args{
				version1: "123",
				version2: "321",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareVersion(tt.args.version1, tt.args.version2); got != tt.want {
				t.Errorf("compareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
