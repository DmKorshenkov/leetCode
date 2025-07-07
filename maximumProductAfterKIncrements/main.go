type Heap []int

func (h *Heap) Push(n any) {
	*h = append(*h, n.(int))
}

func (h *Heap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}


func maximumProduct(nums []int, k int) int {
    h := &Heap{}
    heap.Init(h)
    for i := range nums{
        heap.Push(h, nums[i])
    }

    for i := range k{
        (*h)[0]++
        heap.Fix(h,0)
        _=i
    }
    res := (*h)[0]
    for i := 1; i < h.Len(); i++ {
        res *= (*h)[i]
        res %= 1000000007
    }
    return res
}
