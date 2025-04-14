package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	head := &ListNode{
		Val: 5, Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 3, Next: &ListNode{
					Val: 13, Next: &ListNode{
						Val: 3, Next: &ListNode{
							Val: 5, Next: &ListNode{
								Val: 8, Next: nil}}}}}}}

	removeNodes2(head)

}

func removeNodes2(head *ListNode) *ListNode {

	dummy := &ListNode{}

	for head != nil {
		next := head.Next
		head.Next = dummy
		dummy = head
		head = next
	}

	head = dummy
	dummy = &ListNode{}
	newHead := dummy

	max := head.Val
	for head != nil {
		if head.Val >= max {
			max = head.Val
			dummy.Next = head
			dummy = dummy.Next
		}
		head = head.Next
	}
	dummy.Next = nil

	newHead, _ = revers(newHead.Next)
	return newHead
}

func removeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{}
	newHead := dummy
	//	prev := revers(head)
	for prev, max := revers(head); prev != nil; prev = prev.Next {
		if prev.Val >= max {
			max = prev.Val
			//	fmt.Println(dummy, dummy.Next)
			dummy.Next = prev
			dummy = dummy.Next

		}
	}
	dummy.Next = nil
	newHead, _ = revers(newHead.Next)

	return newHead
}

func revers(head *ListNode) (*ListNode, int) {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev, prev.Val
}
