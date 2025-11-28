package leetcode_test

import (
	"testing"

	"github.com/thanhnamdk2710/algorithm/practice/leetcode"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  []byte
	}{
		{
			name:  "simple case - hello",
			input: []byte{'h', 'e', 'l', 'l', 'o'},
			want:  []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name:  "simple case - Hannah",
			input: []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			want:  []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
		{
			name:  "single character",
			input: []byte{'a'},
			want:  []byte{'a'},
		},
		{
			name:  "two characters",
			input: []byte{'a', 'b'},
			want:  []byte{'b', 'a'},
		},
		{
			name:  "empty array",
			input: []byte{},
			want:  []byte{},
		},
		{
			name:  "palindrome",
			input: []byte{'r', 'a', 'c', 'e', 'c', 'a', 'r'},
			want:  []byte{'r', 'a', 'c', 'e', 'c', 'a', 'r'},
		},
		{
			name:  "numbers and letters",
			input: []byte{'1', '2', 'a', 'b'},
			want:  []byte{'b', 'a', '2', '1'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy since the function modifies in-place
			input := make([]byte, len(tt.input))
			copy(input, tt.input)
			
			got := leetcode.ReverseString(input)
			if len(got) != len(tt.want) {
				t.Errorf("ReverseString() length = %v, want %v", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("ReverseString(%q) = %q, want %q", tt.input, got, tt.want)
					return
				}
			}
		})
	}
}
