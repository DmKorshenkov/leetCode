package main

func main() {

}

func minOperations(s string) int {
    ln := len(s)
    bs1 := make([]byte, ln)
//    bs2 := make([]byte, ln)

    res1 := 0
    for i := 0; i < ln; i++ {
        if i % 2 == 0 {
            bs1[i]++
        }
        if bs1[i] != s[i] - 48 {
            res1++
        }

    }
    return min(res1, ln-res1)
}
