package LeetCode

import "testing"

func TestListTest(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"first", args{[]int{1, 2, 3}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListTest(tt.args.nums); got != tt.want {
				t.Errorf("ListTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
