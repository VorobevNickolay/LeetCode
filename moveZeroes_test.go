package LeetCode

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"first", args{[]int{0, 1, 0, 3, 12}}, []int{1, 3, 12, 0, 0}},
		{"zero", args{[]int{0}}, []int{0}},
		{"already_sorted", args{[]int{1, 0}}, []int{1, 0}},
		{"without nulls", args{[]int{1}}, []int{1}},
		{"One at the end", args{[]int{0, 0, 0, 1}}, []int{1, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveZeroes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
