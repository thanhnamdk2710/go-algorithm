package leetcode_test

import (
	"testing"

	"github.com/thanhnamdk2710/algorithm/practice/leetcode"
)

func TestTwoSumII(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{
			name:    "example 1 - middle elements",
			numbers: []int{2, 7, 11, 15},
			target:  9,
			want:    []int{1, 2},
		},
		{
			name:    "example 2 - adjacent elements",
			numbers: []int{2, 3, 4},
			target:  6,
			want:    []int{1, 3},
		},
		{
			name:    "example 3 - negative numbers",
			numbers: []int{-1, 0},
			target:  -1,
			want:    []int{1, 2},
		},
		{
			name:    "first and last elements",
			numbers: []int{1, 2, 3, 4, 5},
			target:  6,
			want:    []int{1, 5},
		},
		{
			name:    "two elements only",
			numbers: []int{1, 2},
			target:  3,
			want:    []int{1, 2},
		},
		{
			name:    "large numbers",
			numbers: []int{5, 25, 75},
			target:  100,
			want:    []int{2, 3},
		},
		{
			name:    "with duplicates",
			numbers: []int{1, 2, 3, 3, 4, 5},
			target:  6,
			want:    []int{1, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := leetcode.TwoSumII(tt.numbers, tt.target)
			if len(got) != len(tt.want) {
				t.Errorf("TwoSumII(%v, %d) length = %v, want %v", tt.numbers, tt.target, len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("TwoSumII(%v, %d) = %v, want %v", tt.numbers, tt.target, got, tt.want)
					return
				}
			}
		})
	}
}
