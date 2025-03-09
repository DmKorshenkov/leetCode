func heightChecker(heights []int) int {
  dupl := append(make([]int, 0,len(heights)),heights...)
  answer := 0
  sort.Ints(dupl)

  for in := range dupl{
    if dupl[in] != heights[in]{
        answer++
    }
  }
  return answer
}
