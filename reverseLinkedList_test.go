package LeetCode

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example", args{ArrayToList([]int{1, 2, 3, 4, 5})}, []int{5, 4, 3, 2, 1}},
		{"1-el", args{ArrayToList([]int{1})}, []int{1}},
		{"2-el", args{ArrayToList([]int{1, 2})}, []int{2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
