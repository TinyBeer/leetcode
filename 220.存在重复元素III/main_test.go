package main

import "testing"

func Test_containsNearbyAlmostDuplicate(t *testing.T) {
	type args struct {
		nums      []int
		indexDiff int
		valueDiff int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nums = [1,2,3,1], indexDiff = 3, valueDiff = 0",
			args: args{
				nums:      []int{1, 2, 3, 1},
				indexDiff: 3,
				valueDiff: 0,
			},
			want: true,
		},
		{
			name: "nums = [1,5,9,1,5,9], indexDiff = 2, valueDiff = 3",
			args: args{
				nums:      []int{1, 5, 9, 1, 5, 9},
				indexDiff: 2,
				valueDiff: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyAlmostDuplicate(tt.args.nums, tt.args.indexDiff, tt.args.valueDiff); got != tt.want {
				t.Errorf("containsNearbyAlmostDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
