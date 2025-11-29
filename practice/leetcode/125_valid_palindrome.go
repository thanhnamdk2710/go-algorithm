package leetcode

// Problem:
// Given a string, check whether it is a palindrome after converting all characters
// to lowercase and removing all non-alphanumeric characters. An empty string or a
// string with only non-alphanumeric characters is considerd a valid palindrome

// Approach (Two Pointer):
// 1. Initialize left at the beginning and right at the end of the string
// 2. Move the left pointer forward until it points to an alphanumeric character
// 3. Move the right pointer backward until it points to an alphanumeric character
// 4. Compare the lowercase versions of both characters:
// - If they differ, return false
// - If they match, move both pointers inward.
// 5. Continue until left >= right, if no mismatches are found, return true

// Complexity:
// Time: O(n)
// Space: O(1)

func IsPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlphaNum(s[left]) {
			left++
		}
		for left < right && !isAlphaNum(s[right]) {
			right--
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphaNum(b byte) bool {
	return (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') ||
		(b >= '0' && b <= '9')
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
