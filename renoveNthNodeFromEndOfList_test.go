package LeetCode

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Normal one", args{ArrayToList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), 5}, []int{1, 2, 3, 4, 5, 7, 8, 9, 10}},
		{"1-el", args{ArrayToList([]int{1}), 1}, []int{}},
		{"n=1", args{ArrayToList([]int{1, 2, 3}), 1}, []int{1, 2}},
		{"n=len of list", args{ArrayToList([]int{1, 2, 3}), 3}, []int{2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
