package main

func main(){

}

func checkZeroOnes(s string) bool {
    segment0, segment1 := 0,0

    for i := 0; i < len(s); i++ {
        j := i
        for j < len(s) && s[j] == '1' {
            j++
        }
        segment1 = max(j-i+1, segment1)
        i = j
    }

    for i := 0; i < len(s); i++ {
        j := i
        for j < len(s) && s[j] == '0' {
            j++
        }
        segment0 = max(j-i+1, segment0)
        i = j
    }
    return  segment1 > segment0
}
