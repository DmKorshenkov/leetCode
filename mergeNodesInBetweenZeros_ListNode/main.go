func mergeNodes(head *ListNode) *ListNode {
   dummy := &ListNode{}
   newHead := dummy
    var sum int

    for head != nil {
        if sum + head.Val == sum{
            if sum != 0 {
                dummy.Next = &ListNode{Val: sum,}
                dummy = dummy.Next
            }
            sum = 0
        }
        sum+=head.Val
        head = head.Next
    }
   return newHead.Next
}
