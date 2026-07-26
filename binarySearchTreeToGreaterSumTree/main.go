/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
    var inOrder func(*TreeNode)

    sum := 0
    inOrder = func(head *TreeNode) {
        if head == nil {
            return
        }

        inOrder(head.Right)
        sum += head.Val
        head.Val = sum
        inOrder(head.Left)
    }

    inOrder(root)
    return root
}
