package LeetCode

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int
	}{
		{"Two", "II", 2},
		{"Four", "IV", 4},
		{"58", "LVIII", 58},
		{"1994", "MCMXCIV", 1994},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := romanToInt(test.str)
			if got != test.want {
				t.Errorf("%s: got %d, want %d", test.name, got, test.want)
			}
		})
	}
}
