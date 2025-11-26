package leetcode

// Problem:
// Given a 2D slice where accounts[i][j] represents how much money
// customer i has in bank j, return the maximum total wealth
// among all customers.

// Approach:
// 1. Initialize a variable maxWealth = 0
// 2. Loop through each customer:
//    - Calculate the sum of all their bank balances
// 3. Update maxWealth if this sum is greater
// 4. Return maxWealth

// Complexity:
// Time:  O(m * n) — must traverse all balances
// Space: O(1) — only storing running totals

// Edge cases:
// - accounts is empty → return 0
// - a customer has no bank accounts → sum = 0

func MaximumWealth(accounts [][]int) int {
	if len(accounts) == 0 {
		return 0
	}

	maxWalth := 0
	for _, customer := range accounts {
		sum := 0
		for _, money := range customer {
			sum += money
		}
		maxWalth = max(maxWalth, sum)
	}
	return maxWalth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
