package leetcode_test

import (
	"testing"

	"github.com/thanhnamdk2710/algorithm/practice/leetcode"
)

// Helper function to create a linked list from a slice
func createList(values []int) *leetcode.ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &leetcode.ListNode{Val: values[0]}
	current := head
	for i := 1; i < len(values); i++ {
		current.Next = &leetcode.ListNode{Val: values[i]}
		current = current.Next
	}
	return head
}

// Helper function to get the nth node from a list
func getNthNode(head *leetcode.ListNode, n int) *leetcode.ListNode {
	current := head
	for i := 0; i < n && current != nil; i++ {
		current = current.Next
	}
	return current
}

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		name       string
		values     []int
		middleIdx  int // index of the expected middle node
	}{
		{
			name:      "example 1 - odd length list [1,2,3,4,5]",
			values:    []int{1, 2, 3, 4, 5},
			middleIdx: 2, // node with value 3
		},
		{
			name:      "example 2 - even length list [1,2,3,4,5,6]",
			values:    []int{1, 2, 3, 4, 5, 6},
			middleIdx: 3, // node with value 4 (second middle)
		},
		{
			name:      "single node",
			values:    []int{1},
			middleIdx: 0,
		},
		{
			name:      "two nodes",
			values:    []int{1, 2},
			middleIdx: 1, // second node
		},
		{
			name:      "three nodes",
			values:    []int{1, 2, 3},
			middleIdx: 1, // middle node
		},
		{
			name:      "four nodes",
			values:    []int{1, 2, 3, 4},
			middleIdx: 2, // second middle node
		},
		{
			name:      "seven nodes",
			values:    []int{1, 2, 3, 4, 5, 6, 7},
			middleIdx: 3, // middle node
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.values)
			want := getNthNode(head, tt.middleIdx)
			got := leetcode.MiddleNode(head)
			
			if got != want {
				gotVal := -1
				wantVal := -1
				if got != nil {
					gotVal = got.Val
				}
				if want != nil {
					wantVal = want.Val
				}
				t.Errorf("MiddleNode() = node with value %v, want node with value %v", gotVal, wantVal)
			}
		})
	}
}
