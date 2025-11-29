package leetcode

// Problem:
// Given a string, reverse only the vowels ('a', 'e', 'i', 'o', 'u', case-insensitive)
// and return the resulting string. Non-vowel characters must remain in their original positions.

// Approach (Two Pointer):
// 1. Convert the string to a mutable slice
// 2. Initialize two pointers: left at the start, right at the end
// 3. Move left forward until it points to a vowel
// 4. Move right backward until it points to a vowel
// 5. When both pointers point to vowels, swap them and move both pointers inward
// 6. Continue until left >= right
// This ensure only vowels are reversed while all other characters remain unchanged

// Complexity:
// Time: O(n)
// Space: O(n)

func ReverseVowels(s string) string {
	b := []byte(s)

	left, right := 0, len(b)-1
	for left < right {
		for left < right && !isVowel(b[left]) {
			left++
		}

		for left < right && !isVowel(b[right]) {
			right--
		}

		b[left], b[right] = b[right], b[left]

		left++
		right--
	}

	return string(b)
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}

	return false
}
