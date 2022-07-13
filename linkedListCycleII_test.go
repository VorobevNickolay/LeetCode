package LeetCode

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{ListWithCycle([]int{3, 2, 0, -4}, 1)}, 2},
		{"Example 2", args{ListWithCycle([]int{3, 2, 0, -4}, -1)}, -1},
		{"One Node", args{ListWithCycle([]int{1}, -1)}, -1},
		{"Small one", args{ListWithCycle([]int{1, 2}, 0)}, 1},
		{"Empty", args{ListWithCycle([]int{}, -1)}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
