package leetcode

// Problem:
// Given string s and t, determine whether s is a subsequence of t.
// A subsequence is formed by deleting zero or more characters from the original
// string with out changing the order of the remaining characters

// Approach (Two pointers):
// 1. Initialize two pointers i = 0 (for s) and j = 0 (for t)
// 2. Traverse t from left to right:
// - If s[i] matches t[j], move i forward
// - Alway move j forward
// 3. If pointer i reaches the end of s, then every character of s appears in t in the correct order, so return true

// Complexity:
// Time: O(m + n)
// Space: O(1)

func IsSubsequence(s, t string) bool {
	i, j := 0, 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}
