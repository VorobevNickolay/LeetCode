package LeetCode

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example1", args{[]byte("hello")}, "olleh"},
		{"One symbol", args{[]byte("h")}, "h"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseString(tt.args.s); got != tt.want {
				t.Errorf("reverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
