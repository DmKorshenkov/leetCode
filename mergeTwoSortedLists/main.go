package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	ptr1 := &ListNode{
		Val:  1,
		Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}},
	}
	ptr2 := &ListNode{
		Val:  1,
		Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}},
	}
	mergeTwoLists(ptr1, ptr2)

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	arr := make([]int8, 0, 50*2)
	for list1 != nil {
		arr = append(arr, int8(list1.Val))
		list1 = list1.Next
	}
	for list2 != nil {
		arr = append(arr, int8(list2.Val))
		list2 = list2.Next
	}
	if len(arr) == 0 {
		return nil
	}
	for now := 0; now < len(arr); now++ {
		for next := 0; next < len(arr); next++ {
			if arr[now] < arr[next] {
				arr[now], arr[next] = arr[next], arr[now]
			}
		}
	}
	list := new(ListNode)
	ptr := list
	for in := 0; in < len(arr); in++ {
		list.Val = int(arr[in])
		if in == len(arr)-1 {
			break
		}
		list.Next = new(ListNode)
		list = list.Next
	}
	return ptr
}
