type NumArray struct {
    Val []int
}


func Constructor(nums []int) NumArray {
    var sum = make([]int, len(nums)+1)
    for i := 1; i <= len(nums); i++ {
        sum[i] = sum[i-1] + nums[i-1]
    }
    return NumArray{Val: sum}
}


func (this *NumArray) SumRange(left int, right int) int {
    return this.Val[right+1] - this.Val[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
