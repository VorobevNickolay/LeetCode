package LeetCode

import "testing"

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Even", args{ArrayToList([]int{3, 2, 0, -4})}, 0},
		{"Odd", args{ArrayToList([]int{3, 2, 6, -4, 5})}, 6},
		{"1-el", args{ArrayToList([]int{2})}, 2},
		{"Empty", args{ArrayToList([]int{})}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); got != tt.want {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
