package leetcode

import "strconv"

// Problem:
// For each integer from 1 to n, produce:
// - "FizzBuzz" if divisible by both 3 and 5
// - "Fizz" if divisible by 3
// - "Buzz" if divisible by 5
// - otherwise the number itself as a string.

// Approach:
// 1. Create a result slice of size n
// 2. Loop i from 1 to n
// 3. Check divisibility in correct priority
// 4. Assign appropriate string into result[i-1]

// Complexity:
// Time:  O(n)
// Space: O(n)

func FizzBuzz(n int) []string {
	result := make([]string, n)

	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result[i-1] = "FizzBuzz"
		case i%3 == 0:
			result[i-1] = "Fizz"
		case i%5 == 0:
			result[i-1] = "Buzz"
		default:
			result[i-1] = strconv.Itoa(i)
		}
	}
	return result
}
