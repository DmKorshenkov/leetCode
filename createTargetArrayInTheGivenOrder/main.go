type listNode struct {
    Head *node
}
type node struct {
	Val  int
	Next *node
}


func (list *listNode)insert(val int, idx int) {
    if idx == 0 {
        list.Head = &node{Val: val, Next: list.Head}
        return
    }

    curr := list.Head
    var prev *node
    for step := 0; curr != nil && step < idx; step++ {
        prev = curr
        curr = curr.Next
    }
    prev.Next = &node{Val: val, Next: curr}
}


func createTargetArray(nums []int, index []int) []int {
    target := make([]int, 0, len(nums))

    var list = &listNode{Head: &node{Val: nums[0]}}
    for i := 1; i < len(nums); i++ {
        list.insert(nums[i], index[i])
    }

    for list.Head != nil {
        target = append(target, list.Head.Val)
        list.Head = list.Head.Next 
    }

    return target
}


