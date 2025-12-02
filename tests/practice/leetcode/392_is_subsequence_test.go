package leetcode_test

import (
	"testing"

	"github.com/thanhnamdk2710/algorithm/practice/leetcode"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "basic true",
			s:    "abc",
			t:    "ahbgdc",
			want: true,
		},
		{
			name: "basic false",
			s:    "axc",
			t:    "ahbgdc",
			want: false,
		},
		{
			name: "empty s is always subsequence",
			s:    "",
			t:    "ahbgdc",
			want: true,
		},
		{
			name: "both empty strings",
			s:    "",
			t:    "",
			want: true,
		},
		{
			name: "s longer than t",
			s:    "abcd",
			t:    "abc",
			want: false,
		},
		{
			name: "characters out of order",
			s:    "abc",
			t:    "acb",
			want: false,
		},
		{
			name: "repeated characters in t",
			s:    "aaa",
			t:    "aaabaaa",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := leetcode.IsSubsequence(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("IsSubsequence(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
