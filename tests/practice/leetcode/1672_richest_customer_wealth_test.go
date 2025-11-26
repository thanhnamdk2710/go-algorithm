package leetcode_test

import (
	"testing"

	"github.com/thanhnamdk2710/algorithm/practice/leetcode"
)

func TestMaximumWealth(t *testing.T) {
	tests := []struct {
		name     string
		accounts [][]int
		want     int
	}{
		{
			name:     "example 1 - basic case",
			accounts: [][]int{{1, 2, 3}, {3, 2, 1}},
			want:     6,
		},
		{
			name:     "example 2 - different wealth",
			accounts: [][]int{{1, 5}, {7, 3}, {3, 5}},
			want:     10,
		},
		{
			name:     "example 3 - larger numbers",
			accounts: [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}},
			want:     17,
		},
		{
			name:     "single customer single bank",
			accounts: [][]int{{100}},
			want:     100,
		},
		{
			name:     "single customer multiple banks",
			accounts: [][]int{{1, 2, 3, 4, 5}},
			want:     15,
		},
		{
			name:     "multiple customers with zeros",
			accounts: [][]int{{0, 0, 0}, {1, 2, 3}},
			want:     6,
		},
		{
			name:     "all zeros",
			accounts: [][]int{{0, 0}, {0, 0}},
			want:     0,
		},
		{
			name:     "empty accounts",
			accounts: [][]int{},
			want:     0,
		},
		{
			name:     "same wealth for all customers",
			accounts: [][]int{{5, 5}, {5, 5}, {5, 5}},
			want:     10,
		},
		{
			name:     "first customer is richest",
			accounts: [][]int{{10, 10, 10}, {1, 1, 1}, {2, 2, 2}},
			want:     30,
		},
		{
			name:     "last customer is richest",
			accounts: [][]int{{1, 1, 1}, {2, 2, 2}, {10, 10, 10}},
			want:     30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := leetcode.MaximumWealth(tt.accounts)
			if got != tt.want {
				t.Errorf("MaximumWealth() = %v, want %v", got, tt.want)
			}
		})
	}
}
