package LeetCode

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Leetcode #1", args{3, 7}, 28},
		{"Leetcode #2", args{3, 2}, 3},
		{"m,n=1", args{1, 1}, 1},
		{"m,n=10", args{10, 10}, 48620},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
