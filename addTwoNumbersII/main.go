/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    l1 = reverseList(l1)
    l2 = reverseList(l2)
    n := 0

    list := &ListNode{}
    prev := list
    for l1 != nil || l2 != nil {

        if l1 != nil {
            n += l1.Val
            l1 = l1.Next
        }

        if l2 != nil {
            n += l2.Val
            l2 = l2.Next
        }

        list.Next = &ListNode{Val: n % 10}
        list = list.Next
        n /= 10
    }
    if n != 0 {
        list.Next =&ListNode{Val: n}
    }
//    prev := reverseList(prev.Next)
    return reverseList(prev.Next)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
