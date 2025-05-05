func getConcatenation(nums []int) []int {
   res := make([]int,0,len(nums)+len(nums))
   res = append(res, nums...)
   res = append(res, nums...)

    return res
//    runtime.GC()

//   return append(append([]int{}, nums...), nums...) 
}
