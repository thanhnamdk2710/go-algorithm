package leetcode

// Problem:
// Count how many operations are needed to reduce num to 0
// - If num is even → divide by 2
// - If num is odd → subtract 1

// Approach:
// 1. Initialize step = 0
// 2. While num > 0
// - If even: num /= 2
// - Else: num--
// - step++
// 3. Return step

// Complexity:
// Time:  O(n)
// Space: O(n)

// Edge case:
// num = 0 → return 0

func NumberOfSteps(num int) int {
	if num == 0 {
		return 0
	}

	step := 0
	for num > 0 {
		if num&1 == 0 { // num/2 == 0
			num >>= 1 // num /= 2
		} else {
			num--
		}
		step++
	}

	return step
}
