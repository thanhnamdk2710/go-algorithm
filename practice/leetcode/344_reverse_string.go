package leetcode

// Problem:
// Given an array of characters, reverse the array in-place
// You must modify the array directly using O(1) extra space

// Approach:
// Use the two pointers technique:
// - Initialize two indices: left at the start and right at the end of the array
// - While left < right:
//   - Swap the characters at positions left and right
//   - Move left forward and right backward

// Complexity:
// Time: O(n)
// Space: O(1)

func ReverseString(s []byte) []byte {
	left, right := 0, len(s)-1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	return s
}
