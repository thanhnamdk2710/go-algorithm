# Understanding the Two Pointers Pattern

When you start solving array and string problems, one pattern appears everywhere: Two Pointers. It turns many O(n²) brute-force solutions into clean, linear time O(n) algorithms.

In this article, we'll cover:

- What the two pointers technique is
- The main patterns (opposite ends, same direction, fast/slow)
- LeetCode-style example
- A realistic backend example you might actually use in production

## 1. What is the Two Pointers technique?

The Two Pointers technique uses two indices (or "pointers") to traverse a collection:

- Both on the same array/string
- Sometimes on two different arrays
- They can move:
  - from opposite ends toward the middle
  - in the same direction but at different speeds

You typically use two pointers when:

- The data is sorted
- You only need to scan once or twice
- You want to reduce complexity from O(n²) → O(n)

### Why it matters

Instead of checking all pairs with two nested loops:

```go
for i := 0; i < n; i++ {
    for j := i + 1; j < n; j++ {
        // O(n²)
    }
}
```

you can often do it with one pass and two pointers:

```go
i, j := 0, n-1

for i < j {
    // O(n)
}
```

This is the core idea

## 2. Core Two Pointers Patterns

Let's break it down into three useful patterns.

### 2.1 Opposite Ends (Left & Right Pointers)

You put one pointer at the start (`left := 0`) and one at the end (`right := len(nums)-1`) and move them toward each other.

Typical use cases:

- Reverse an array/string
- Check if a string is a palindrome
- Container With Most Water
- Two Sum on a sorted array

Example: reverse an array in-place

```go
func reverse(nums []int) {
    left, right := 0, len(nums)-1
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
}
```

Time complexity: O(n)
Space complexity: O(1)

### 2.2 Same Direction (Sliding-Style Two Pointers)

Both pointers move left → right:

- `slow` tracks the position to write or window start
- `fast` scans through the array / string

Typical use cases:

- Remove duplicates from sorted array
- Move zeroes to the end
- Compress in-place
- Build sliding window

Example: remove duplicates from a sorted array

```go
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    write := 1
    for read := 1; read < len(nums); read++ {
        if nums[read] != nums[read-1] {
            nums[write] = nums[read]
            write++
        }
    }
    return write // new length
}
```

### 2.3 Fast & Slow Pointers

Two pointers move in the same direction, but one is faster:

- `slow` moves by 1 step
- `fast` moves by 2 steps

Typical use cases:

- Detect cycle in Linked List
- Find middle node in Linked List
- Detect cycle in sequence/state transitions

Even though it's often used on Linked Lists, the mental model is still "two pointers with different speed".

## 3. LeetCode Example

### 3.1. Valid Palindrome (Opposite Ends)

**Prolem (simplified):**

Given a string `s`, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

This is almost textbook **Two Pointer**.

**Brute force idea**

- Clean string to only alphanumerics → `filtered`
- Reverse `filtered` and compare with original

This works but uses extra memory and additional passes.

**Two Pointers approach**

- Use two indices: `left` at start, `right` at end
- Move inward
- Skip non-alphanumeric characters
- Compare lowercase versions

Implementation:

```go
import "unicode"

func isPalindrome(s string) bool {
    left, right := 0, len(s)-1

    for left < right {
        // move left to next alphanumeric
        for left < right && !isAlnum(rune(s[left])) {
            left++
        }

        // move right to previous alphanumeric
        for left < right && !isAlnum(rune(s[right])) {
            right--
        }

        if left < right {
            if toLower(rune(s[left])) != toLower(rune(s[right])) {
                return false
            }
            left++
            right--
        }
    }
    return true
}

func isAlnum(r rune) bool {
    return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func toLower(r rune) rune {
    return unicode.ToLower(r)
}
```

Key points:

- Two pointers from both ends.
- Skip invalid characters inside the loop (no extra arrays).
- Time: O(n), Space: O(1) (not counting the string itself).

### 3.2. Two Sum II - Input Array Is Sorted

**Problem:**

Given a sorted array of integers `nums` and a target integer `target`, return the 1-based indices of the two numbers that add up to `target`.

The sorted property is key: perfect for two pointers.

**Idea**

- Use `left` at the beginning, `right` at the end
- If `nums[left] + nums[right]` is:
  - `< target` → move `left++`
  - `> target` → move `right--`
  - `== target` → found

Implementation:

```go
func twoSumSorted(nums []int, target int) []int {
    left, right := 0, len(nums)-1

    for left < right {
        sum := nums[left] + nums[right]
        if sum == target {
            // return 1-based indices
            return []int{left + 1, right + 1}
        } else if sum < target {
            left++
        } else {
            right--
        }
    }

    return nil
}
```

Why this is better than brute force:

- Brute force: nested loops → O(n²)
- Two pointers: single scan from both ends → O(n)

This pattern generalizes to problems like:

- 3Sum (sort + fix one index + two pointers)
- Container With Most Water
- Pair sum closest to X

## 4. Backend Example: Merging Sorted Event Streams

Two pointers aren't just for competitive programming. They show up in backend systems when you work with sorted data, time-series, or logs.

**Scenario**

Imagine you have two event sources:

- Service A: `[]Event` sorted by `Timestamp`
- Service B: `[]Event` sorted by `Timestamp`

You want to merge them into a single, time-ordered stream for:

- Analytics
- Building a unified audit log
- Debugging a complex incident

This is basically the merge step of Merge Sort and is a classic two pointers problem.

Data model:

```go
type Event struct {
    Source string
    Timestamp int64
    Message string
}
```

Two Pointers solution:

```go
func mergeEvents(a, b []Event) []Event {
    i, j := 0, 0
    result := make([]Event, len(a)+len(b))

    for i < len(a) && j < len(b) {
        if a[i].Timestamp <= b[j].Timestamp {
            result = append(result, a[i])
            i++
        } else {
            result = append(result, b[j])
            j++
        }
    }

    // Append remaining events
    for i < len(a) {
        result = append(result, a[i])
        i++
    }
    for j < len(b) {
        result = append(result, b[j])
        j++
    }

    return result
}
```

Why this works well:

- Both arrays are already sorted → no need to sort the combined array
- Time complexity: O(n + m) when n, m are lengths of the two arrays
- Space complexity: O(n + m) for the result slice

Compare that with:

- Concatenate the sort: O((n + m) log(n + m))

In systems with large logs/event streams, this difference is huge.

## 5. Another Practical Use: Deduplicating Sorted IDs

You might have a batch job that:

- Reads user IDs from a file or database
- The IDs are sorted
- You want to remove duplicates before processing

Instead of using a `map[int]bool` and extra memory, two pointers let you deduplicate in-place.

```go
func uniqueSortedIDs(ids []int) []int {
    if len(ids) == 0 {
        return ids
    }

    write := 1
    for read := 1; read < len(ids); read++ {
        if ids[read] != ids[read-1] {
            ids[write] = ids[read]
            write++
        }
    }

    return ids[:write] // slice with unique IDs
}
```

This is a production-ready pattern when:

- The data is large
- Memory matters
- Data is already sorted (e.g., from DB `ORDER BY`)

## 6. Common Pitfalls and Best Practices

### 6.1. Off-by-one errors

Typical bugs:

- Using `<=` instead `<`
- Accessing `nums[i+1]` without checking `i+1 < len(nums)`

Always double-check loop conditions when using two pointers.

### 6.2. Strings: bytes vs runes

For ASCII-only problems, using `s[i] // byte` is fine.

For Unicode (Vietnamese, emoji, etc.), you must convert to `[]rune`:

```go
r := []rune(s)
left, right := 0, len(r)-1
// work on r[left], r[right]
```

Otherwise, you risk splitting multi-byte characters incorrectly.

### 6.3. Keep the invariant clear

Every two pointers algorithm has a loop invariant:

A condition that is always true inside the loop

Example:

- Palindrome check: "All characters between original bounds and current pointers have been checked and are valid."
- Two Sum sorted: "All pairs that could be a solution are between `left` and `right`."

When writing complex logic, write down the invariant in a comment first. It makes reasoning easier.

## 7. Summary

The Two Pointers technique is one of the most powerful tools for array and string problems:

- It converts many quadratic algorithms into linear ones.
- It appears in interview problems
- It works especially well with:
  - Sorted arrays
  - Palindrome
  - Sliding windows
  - Merging and deduplicating sequences

If you're practicing algorithms, you should:

1. Learn to recognize when two pointers apply.
2. Implement common patterns:
   - Opposite ends (palindrome, container)
   - Same direction (compress, remove duplicates)
   - Fast/slow (cycle detection)
3. Appy the same thinking in your day-to-day backend work:
   - Merging sorted event streams
   - Deduplicating sorted IDs
   - Scanning time-series data
