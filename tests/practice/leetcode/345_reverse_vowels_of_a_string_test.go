package leetcode_test

import (
	"testing"

	"github.com/thanhnamdk2710/algorithm/practice/leetcode"
)

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example 1 - hello",
			input: "hello",
			want:  "holle",
		},
		{
			name:  "example 2 - leetcode",
			input: "leetcode",
			want:  "leotcede",
		},
		{
			name:  "single vowel",
			input: "a",
			want:  "a",
		},
		{
			name:  "no vowels",
			input: "bcdfg",
			want:  "bcdfg",
		},
		{
			name:  "all vowels",
			input: "aeiou",
			want:  "uoiea",
		},
		{
			name:  "mixed case vowels",
			input: "AEIOUaeiou",
			want:  "uoieaUOIEA",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "consonants only",
			input: "xyz",
			want:  "xyz",
		},
		{
			name:  "mixed case - Hannah",
			input: "Hannah",
			want:  "Hannah",
		},
		{
			name:  "with spaces and punctuation",
			input: "race car",
			want:  "race car",
		},
		{
			name:  "uppercase vowels",
			input: "HELLO",
			want:  "HOLLE",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := leetcode.ReverseVowels(tt.input)
			if got != tt.want {
				t.Errorf("ReverseVowels(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
