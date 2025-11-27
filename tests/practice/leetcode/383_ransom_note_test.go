package leetcode_test

import (
	"testing"

	"github.com/thanhnamdk2710/algorithm/practice/leetcode"
)

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		name       string
		ransomNote string
		magazine   string
		want       bool
	}{
		{
			name:       "simple case - can construct",
			ransomNote: "a",
			magazine:   "b",
			want:       false,
		},
		{
			name:       "cannot construct - not enough letters",
			ransomNote: "aa",
			magazine:   "ab",
			want:       false,
		},
		{
			name:       "can construct - exact match",
			ransomNote: "aa",
			magazine:   "aab",
			want:       true,
		},
		{
			name:       "empty ransom note",
			ransomNote: "",
			magazine:   "abc",
			want:       true,
		},
		{
			name:       "ransom note longer than magazine",
			ransomNote: "aaaa",
			magazine:   "aa",
			want:       false,
		},
		{
			name:       "can construct - multiple letters",
			ransomNote: "abc",
			magazine:   "aabbcc",
			want:       true,
		},
		{
			name:       "cannot construct - missing letter",
			ransomNote: "xyz",
			magazine:   "abc",
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := leetcode.CanConstruct(tt.ransomNote, tt.magazine)
			if got != tt.want {
				t.Errorf("CanConstruct(%q, %q) = %v, want %v", tt.ransomNote, tt.magazine, got, tt.want)
			}
		})
	}
}
