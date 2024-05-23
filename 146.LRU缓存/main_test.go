package main

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		ops  []string
		para [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{
			name: `["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]`,
			args: args{
				ops:  []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
				para: [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			},
			want: "[null, null, null, 1, null, -1, null, -1, 3, 4]",
		},
	}
	for _, tt := range tests {
		tmp := make([]string, len(tt.args.ops))
		t.Run(tt.name, func(t *testing.T) {
			var cache LRUCache
			for i, op := range tt.args.ops {
				switch op {
				case "LRUCache":
					cache = Constructor(tt.args.para[i][0])
					tmp[i] = "null"
				case "put":
					cache.Put(tt.args.para[i][0], tt.args.para[i][1])
					tmp[i] = "null"
				case "get":
					got := cache.Get(tt.args.para[i][0])
					tmp[i] = strconv.Itoa(got)
				}
			}
			res := "[" + strings.Join(tmp, ", ") + "]"
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("result: %v, want %v", res, tt.want)
			}
		})
	}
}
