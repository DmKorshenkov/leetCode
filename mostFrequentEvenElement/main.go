func mostFrequentEven(nums []int) int {
    evenHash := make(map[int]int)

    for i :=0; i < len(nums);i++ {
        if nums[i]%2 == 0 {
            evenHash[nums[i]]++
        }
    }

    freq := 0
    mostFreqElem := -1
    for n, freqHash := range evenHash {
        if freqHash > freq || freqHash == freq && n < mostFreqElem {
            freq = freqHash
            mostFreqElem = n
        }
    }
    return mostFreqElem
}
