/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    var arr []int
    var post func(root *TreeNode)
    post = func(root *TreeNode) {
       if root == nil {
        return
       } 
       post(root.Left)
       post(root.Right)
       arr = append(arr, root.Val)
    }

    post(root)
    return arr
}
