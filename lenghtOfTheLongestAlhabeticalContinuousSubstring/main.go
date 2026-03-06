package main

func main(){

}

func longestContinuousSubstring(s string) int {
    res, count := 1, 1
    for i := 1; i < len(s); i++ {
        if s[i] - s[i-1] != 1 {
            res = max(count, res)
            count = 0
        }
        count++
    }
    return max(res,count)
}
