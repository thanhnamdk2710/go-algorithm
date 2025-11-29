package leetcode_test

import (
	"testing"

	"github.com/thanhnamdk2710/algorithm/practice/leetcode"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "example 1 - valid palindrome with punctuation",
			input: "A man, a plan, a canal: Panama",
			want:  true,
		},
		{
			name:  "example 2 - not a palindrome",
			input: "race a car",
			want:  false,
		},
		{
			name:  "example 3 - empty string",
			input: " ",
			want:  true,
		},
		{
			name:  "single character",
			input: "a",
			want:  true,
		},
		{
			name:  "simple palindrome",
			input: "racecar",
			want:  true,
		},
		{
			name:  "palindrome with mixed case",
			input: "RaceCar",
			want:  true,
		},
		{
			name:  "palindrome with numbers",
			input: "A1b2B1a",
			want:  true,
		},
		{
			name:  "not palindrome with numbers",
			input: "0P",
			want:  false,
		},
		{
			name:  "only non-alphanumeric characters",
			input: ".,!@#",
			want:  true,
		},
		{
			name:  "empty string",
			input: "",
			want:  true,
		},
		{
			name:  "palindrome with spaces and punctuation",
			input: "Was it a car or a cat I saw?",
			want:  true,
		},
		{
			name:  "not palindrome with spaces",
			input: "hello world",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := leetcode.IsPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
