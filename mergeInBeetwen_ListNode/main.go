package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	mergeInBetween(nil, 1, 3, nil)
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {

	var nodeA *ListNode
	var nodeB *ListNode
	newHead := list1
	for list1 != nil {
		a--
		if a == 0 {
			nodeA = list1
		}
		if b == -1 {
			nodeB = list1
		}
		b--
		list1 = list1.Next
	}

	nodeA.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = nodeB
	return newHead
}
