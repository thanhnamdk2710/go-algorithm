package leetcode_test

import (
	"testing"

	"github.com/thanhnamdk2710/algorithm/practice/leetcode"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example 1 - multiple triplets",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "example 2 - no triplets",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "example 3 - all zeros",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "empty array",
			nums: []int{},
			want: [][]int{},
		},
		{
			name: "array with less than 3 elements",
			nums: []int{1, 2},
			want: [][]int{},
		},
		{
			name: "no valid triplets",
			nums: []int{1, 2, 3},
			want: [][]int{},
		},
		{
			name: "with duplicates - should return unique triplets",
			nums: []int{-2, 0, 0, 2, 2},
			want: [][]int{{-2, 0, 2}},
		},
		{
			name: "larger array with multiple solutions",
			nums: []int{-4, -1, -1, 0, 1, 2, 3},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}, {-4, 1, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := leetcode.ThreeSum(tt.nums)
			if len(got) != len(tt.want) {
				t.Errorf("ThreeSum(%v) returned %d triplets, want %d triplets\nGot: %v\nWant: %v",
					tt.nums, len(got), len(tt.want), got, tt.want)
				return
			}

			// Compare each triplet
			for i := range got {
				if len(got[i]) != 3 {
					t.Errorf("ThreeSum(%v) triplet[%d] has length %d, want 3", tt.nums, i, len(got[i]))
					return
				}

				// Check if triplet matches expected
				match := false
				for j := range tt.want {
					if slicesEqual(got[i], tt.want[j]) {
						match = true
						break
					}
				}
				if !match {
					t.Errorf("ThreeSum(%v) = %v, want %v", tt.nums, got, tt.want)
					return
				}
			}
		})
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
