package LeetCode

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example", args{"aba", "aab"}, true},
		{"False", args{"aaa", "aabaa"}, false},
		{"1 el s2", args{"abc", "a"}, false},
		{"Leetcode #1", args{"ab", "eidbaooo"}, true},
		{"Leetcode #2", args{"ab", "eidboao"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
