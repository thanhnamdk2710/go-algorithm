package leetcode_test

import (
	"testing"

	"github.com/thanhnamdk2710/algorithm/practice/leetcode"
)

func TestNumberOfSteps(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{
			name: "example 1: num = 14",
			num:  14,
			want: 6,
		},
		{
			name: "example 2: num = 8",
			num:  8,
			want: 4,
		},
		{
			name: "example 3: num = 123",
			num:  123,
			want: 12,
		},
		{
			name: "edge case: num = 0",
			num:  0,
			want: 0,
		},
		{
			name: "edge case: num = 1",
			num:  1,
			want: 1,
		},
		{
			name: "power of 2: num = 16",
			num:  16,
			want: 5,
		},
		{
			name: "power of 2: num = 32",
			num:  32,
			want: 6,
		},
		{
			name: "odd number: num = 7",
			num:  7,
			want: 5,
		},
		{
			name: "large even: num = 100",
			num:  100,
			want: 9,
		},
		{
			name: "large odd: num = 99",
			num:  99,
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := leetcode.NumberOfSteps(tt.num)
			if got != tt.want {
				t.Errorf("NumberOfSteps(%d) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
