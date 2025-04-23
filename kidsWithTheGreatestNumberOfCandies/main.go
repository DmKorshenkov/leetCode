func kidsWithCandies(candies []int, extraCandies int) []bool {
    max := slices.Max(candies)
    res := []bool{}
    for i := range candies {
        res = append(res, candies[i] + extraCandies >= max)
    }
    return res
}
