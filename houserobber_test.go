package LeetCode

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Leetcode #1", args{[]int{1, 2, 3, 1}}, 4},
		{"Leetcode #2", args{[]int{2, 7, 9, 3, 1}}, 12},
		{"Len(nums)=1", args{[]int{1}}, 1},
		{"Len(nums)=2", args{[]int{1, 2}}, 2},
		{"Len(nums)=3", args{[]int{1, 2, 4}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
