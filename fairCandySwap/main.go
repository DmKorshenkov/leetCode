func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
    var aliceSum int
    var bobSum int
    for i := range aliceSizes{
        aliceSum += aliceSizes[i]
    }

    for i := range bobSizes{
        bobSum += bobSizes[i]
    }

    for i := 0; i < len(aliceSizes); i++{
        aGift := aliceSizes[i]
        for j := 0; j < len(bobSizes); j++{
            
            if aliceSum - aGift + bobSizes[j] == bobSum - bobSizes[j] + aGift {
                return []int{aGift, bobSizes[j]}
            } 
        }
    }

    return nil
}
