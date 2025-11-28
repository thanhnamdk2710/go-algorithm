package leetcode

// Problem:
// Given a sorted (non-decreasing) array of integers, return the 1-indexed positions of the two numbers whose sum equals the target
// Exactly one solution exists, and each number may be used only once
// The solution must use O(1) extra space

// Approach (Two-pointer):
// 1. Initialize two indices: left at the start and right at the end of the array
// 2. While left < right:
//	 - Compute the sum of numbers[left] + numbers[right]
//   - If the sum equals the target, return the 1-indexed positions.
//   - If the sum is greater than the target, move the right pointer leftward to decrease the sum.
//   - If the sum is less than the target, move the left pointer rightward to increase the sum.
// This works because the array is already sorted.

// Complexity:
// Time: O(n)
// Space: O(1)

func TwoSumII(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}
		if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}

	return nil
}
