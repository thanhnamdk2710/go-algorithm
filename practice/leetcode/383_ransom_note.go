package leetcode

// Problem:
// Check if ransomNote can be constructed using characters from magazine
// Each character in magazine can only be used once

// Approach:
// 1. Create a frequency array freq[26] for lowercase letters
// 2. Count occurrences of each letter in magazine
// 3. Loop through ransomNote and decrement freq for each character
// 4. If any count goes below zero → magazine does not have enough letters
// 5. Otherwise return true

// Why array[26]?
// - Characters are guaranteed to be lowercase 'a' to 'z'
// - Array lookup is faster than map
// - Memory cost is constant

// Complexity:
// Time: O(m + n)
// Space: O(1)

func CanConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	freq := [26]int{}

	for _, ch := range magazine {
		freq[ch-'a']++
	}

	for _, ch := range ransomNote {
		freq[ch-'a']--
		if freq[ch-'a'] < 0 {
			return false
		}
	}

	return true
}
