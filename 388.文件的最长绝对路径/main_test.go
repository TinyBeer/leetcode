package main

import "testing"

func Test_lengthLongestPath(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `input = "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"`,
			args: args{
				input: "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext",
			},
			want: 20,
		},
		{
			name: `input = "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"`,
			args: args{
				input: "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext",
			},
			want: 32,
		},
		{
			name: `input = "a"`,
			args: args{
				input: "a",
			},
			want: 0,
		},
		{
			name: `input = "file1.txt\nfile2.txt\nlongfile.txt"`,
			args: args{
				input: "file1.txt\nfile2.txt\nlongfile.txt",
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthLongestPath(tt.args.input); got != tt.want {
				t.Errorf("lengthLongestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
