package LeetCode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"easy", args{[]int{1, 2, 3}, 4}, []int{1, 3}},
		{"medium", args{[]int{1, 2, 3, 4, 5, 6, 8}, 14}, []int{6, 7}},
		{"hard", args{[]int{1, 2, 3, 4, 7, 8}, 7}, []int{3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
