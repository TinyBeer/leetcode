package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "head = [1,2,3,4,5], n = 2",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
				n: 2,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
		{
			name: "head = [1], n = 1",
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
				n: 1,
			},
			want: nil,
		},
		{
			name: "head = [1,2], n = 1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
				n: 1,
			},
			want: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeNthFromEnd(tt.args.head, tt.args.n)
			gjs, _ := json.Marshal(got)
			wjs, _ := json.Marshal(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %s, want %s", gjs, wjs)
			}
		})
	}
}
