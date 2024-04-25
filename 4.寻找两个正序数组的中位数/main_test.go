package main

import (
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "nums1 = [1,3], nums2 = [2]",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2,
		},
		{
			name: "nums1 = [1,2], nums2 = [3,4]",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.5,
		},

		{
			name: "nums1 = [0,0], nums2 = [0,0]",
			args: args{
				nums1: []int{0, 0},
				nums2: []int{0, 0},
			},
			want: 0,
		},
		{
			name: "nums1 = [0,0,0,0,0], nums2 = [-1,0,0,0,0,0,1]",
			args: args{
				nums1: []int{0, 0, 0, 0, 0},
				nums2: []int{-1, 0, 0, 0, 0, 0, 1},
			},
			want: 0,
		},

		{
			name: "nums1 = [1,3], nums2 = [2,7]",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2, 7},
			},
			want: 2.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getK(t *testing.T) {
	type args struct {
		num1   []int
		start1 int
		end1   int
		num2   []int
		start2 int
		end2   int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums1 = [1,2], nums2 = [3,4] k = 1",
			args: args{
				num1:   []int{1, 2},
				start1: 0,
				end1:   1,
				num2:   []int{3, 4},
				start2: 0,
				end2:   0,
				k:      1,
			},
			want: 2,
		},
		{
			name: "nums1 = [1,3], nums2 = [2] k = 2",
			args: args{
				num1:   []int{1, 3},
				start1: 0,
				end1:   1,
				num2:   []int{2},
				start2: 0,
				end2:   0,
				k:      2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getK(tt.args.num1, tt.args.start1, tt.args.end1, tt.args.num2, tt.args.start2, tt.args.end2, tt.args.k); got != tt.want {
				t.Errorf("getK() = %v, want %v", got, tt.want)
			}
		})
	}
}
