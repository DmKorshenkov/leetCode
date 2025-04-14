package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {

	if head.Next == nil {
		return head
	}
	newHead := head
	//14(7)->7 7
	for head != nil && head.Next != nil {
		next := head.Next
		newNode := &ListNode{Val: gcd(head.Val, head.Next.Val), Next: next}
		head.Next = newNode
		head = head.Next.Next
	}
	return newHead
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
