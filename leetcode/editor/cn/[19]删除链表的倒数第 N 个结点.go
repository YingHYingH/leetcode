package src

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, head}
	x := findFromEnd(dummy, n+1)
	x.Next = x.Next.Next
	return dummy.Next
}

func findFromEnd(head *ListNode, k int) *ListNode {
	p1 := head
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	p2 := head
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}

//leetcode submit region end(Prohibit modification and deletion)
