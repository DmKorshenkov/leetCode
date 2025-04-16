package main

import (
	"fmt"
)

type ListNode struct {
	Key  int
	Arr  []int
	Next *ListNode
	Up   *ListNode
	Down *ListNode
}

func main() {
	groupThePeople([]int{3, 3, 3, 2, 1, 2})
}

func groupThePeople(groupSizes []int) [][]int {
	current := &ListNode{}
	/*	for i := 0; i < len(groupSizes); i++ {
		current.insert(groupSizes[i], i)
		break
	}*/
	current.insert(3, 1)
	current.insert(3, 2)
	fmt.Println(current)
	return nil
}

func (head *ListNode) insert(key, val int) {
	if head == nil {
		fmt.Println("0")
		return
	}

	return
}
