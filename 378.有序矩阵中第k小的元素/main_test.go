package main

import "testing"

func Test_kthSmallest(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8",
			args: args{
				matrix: [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}},
				k:      8,
			},
			want: 13,
		},
		{
			name: "matrix = [[-5]], k = 1",
			args: args{
				matrix: [][]int{{-5}},
				k:      1,
			},
			want: -5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
