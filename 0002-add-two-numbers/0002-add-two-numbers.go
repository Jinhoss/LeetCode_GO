/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    tempHead := &ListNode{}
    cur, carry := tempHead, 0

    for l1 != nil || l2 != nil {
        val1, val2 := 0, 0
        if l1 != nil {
            val1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            val2 = l2.Val
            l2 = l2.Next
        }

        var sum int = val1 + val2 + carry
        carry = sum / 10

        cur.Next = &ListNode{Val: sum % 10}
        cur = cur.Next 
    
    }

    if carry > 0 {
        cur.Next = &ListNode{Val: carry}
    }
    return tempHead.Next
    
}