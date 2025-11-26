package leetcode_test

import (
	"testing"

	"github.com/thanhnamdk2710/algorithm/practice/leetcode"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "n = 1",
			n:    1,
			want: []string{"1"},
		},
		{
			name: "n = 3 - includes Fizz",
			n:    3,
			want: []string{"1", "2", "Fizz"},
		},
		{
			name: "n = 5 - includes Buzz",
			n:    5,
			want: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name: "n = 15 - includes FizzBuzz",
			n:    15,
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
		{
			name: "n = 20 - comprehensive test",
			n:    20,
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz"},
		},
		{
			name: "n = 30 - multiple FizzBuzz",
			n:    30,
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz", "Fizz", "22", "23", "Fizz", "Buzz", "26", "Fizz", "28", "29", "FizzBuzz"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := leetcode.FizzBuzz(tt.n)
			if len(got) != len(tt.want) {
				t.Errorf("FizzBuzz() length = %v, want %v", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("FizzBuzz()[%d] = %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}
