func stableMountains(height []int, threshold int) []int {
    step := 0
    for i := 1; i < len(height); i++ {
        if height[i-1] > threshold {
            height[step] = i
            step++
        }
    }
    return height[0:step]
}
