package LeetCode

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Normal one", args{"aaaaccccb"}, 9},
		{"With one letter", args{"a"}, 1},
		{"Without middle", args{"aaaacccc"}, 8},
		{"All different letters", args{"qwertyuiop"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
