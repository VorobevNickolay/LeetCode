package LeetCode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Leetcode #1", args{"abcabcbb"}, 3},
		{"Leetcode #2", args{"bbbbb"}, 1},
		{"Leetcode #3", args{"pwwkew"}, 3},
		{"VN #1", args{"aabbaabb"}, 2},
		{"VN #2", args{"a"}, 1},
		{"VN #3", args{"abc"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
