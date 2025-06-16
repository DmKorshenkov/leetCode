func inorderTraversal(root *TreeNode) []int {
    var arr []int
    inorder(root, &arr)
    return arr
}

func inorder(root *TreeNode, arr *[]int) {
    if root == nil {
        return
    }

    inorder(root.Left, arr)
    *arr = append(*arr, root.Val)
    inorder(root.Right, arr)
}
