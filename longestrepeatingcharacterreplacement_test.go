package LeetCode

import "testing"

func Test_characterReplacement(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Leetcode #1", args{"ABAB", 2}, 4},
		{"Leetcode #2", args{"AABABBA", 1}, 4},
		{"k=0", args{"AABABBA", 0}, 2},
		{"useless k", args{"AAAAAAA", 4}, 7},
		{"1-el", args{"A", 1}, 1},
		{"k=len(s)", args{"abcdef", 6}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacement(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("characterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
