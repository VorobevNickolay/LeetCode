package LeetCode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example", args{ArrayToList([]int{0, 1, 0, 1, 0})}, true},
		{"2-el true", args{ArrayToList([]int{0, 0})}, true},
		{"2-el false", args{ArrayToList([]int{0, 1, 1})}, false},
		{"1-el", args{ArrayToList([]int{1})}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
