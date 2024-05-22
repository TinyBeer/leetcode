package main

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "head = [3,2,0,-4], pos = 1",
			args: args{
				head: genList([]int{3, 2, 0, -4}, 1),
			},
			want: true,
		},
		{
			name: "head = [1,2], pos = 0",
			args: args{
				head: genList([]int{1, 2}, 0),
			},
			want: true,
		},
		{
			name: "head = [1], pos = -1",
			args: args{
				head: genList([]int{1}, -1),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
