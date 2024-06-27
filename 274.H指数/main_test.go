package main

import "testing"

func Test_hIndex(t *testing.T) {
	type args struct {
		citations []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "citations = [3,0,6,1,5]",
			args: args{
				citations: []int{3, 0, 6, 1, 5},
			},
			want: 3,
		},
		{
			name: "citations = [1,3,1]",
			args: args{
				citations: []int{1, 3, 1},
			},
			want: 1,
		},
		{
			name: "citations = [100]",
			args: args{
				citations: []int{100},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex(tt.args.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
