package main

func main(){

}

func maxPower(s string) int {
    countC, res := 1, 1

    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1] {
            countC++
        } else {
            res = max(res,countC)
            countC = 1
        }
        res = max(res, countC)
    }

    return res
}
