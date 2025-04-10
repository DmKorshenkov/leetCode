package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode)insert(val int) {
	if t == nil {
		return
	}
	if val < t.Val {
		t.Left
	}

}


func main() {

}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return nil
}
