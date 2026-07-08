func isPalindrome(head *ListNode) bool {
    mid, fast := head, head
    for fast != nil && fast.Next != nil {
        mid = mid.Next
        fast = fast.Next.Next
    }

    var prev *ListNode
    curr := mid
    for curr != nil{
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    for head != nil && prev != nil {
        if head.Val != prev.Val {
            return false
        }
        head = head.Next
        prev = prev.Next
    }
    return true
}
