package leetcode

// Problem:
// Return the middle node of a singly linked list
// If the list has two middle nodes, return the second one

// Approach (Fast-Slow Pointer):
// 1. Initialize two pointers slow and fast at the head
// 2. Move slow by 1 step each iteration
// 3. Move fast by 2 steps each iteration
// 4. When fast reaches the end (nil or fast.Next == nil), slow will be at the middle node
// 5. Return slow

// Why this works:
// - For odd lenght: fast reaches the end first → slow is exactly in the middle
// - For even length: fast reaches nil after jumping over the last node → slow ends at the second middle node as required

// Complexity:
// Time: O(n)
// Space: O(1)

// Edge cases:
// - Empty list → return nil
// - 1 or 2 nodes → logic still works correctly

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
