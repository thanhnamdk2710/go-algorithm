package leetcode

import "sort"

// Problem:
// Given an integer array nums, return all unique triplets [nums[i], nums[j], nums[k]]
// such that their sum is zero. The result must not contain duplicate triplets.

// Approach:
// 1. Sort the array to enable the two-pointer technique and make it possible to skip duplicates.
// 2. Iterate through nums with index i:
//    - Skip nums[i] if it is the same as nums[i-1] to avoid duplicate triplets.
//    - For each i, use two pointers:
//      * left = i + 1
//      * right = end of array
//    - Compute the sum = nums[i] + nums[left] + nums[right].
//    - If sum == 0:
//        • Append the triplet.
//        • Move both pointers inward.
//        • Skip duplicate values of nums[left] and nums[right].
//    - If sum < 0: move left forward to increase sum.
//    - If sum > 0: move right backward to decrease sum.

// Complexity:
// Time: O(n²)
// Space: O(1)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	for i := 0; i < len(nums)-2; i++ {

		// Skip duplicate values for i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				left++
				right--

				// Skip duplicates for left
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				// Skip duplicates for right
				for left < right && nums[right] == nums[right+1] {
					right--
				}

			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
