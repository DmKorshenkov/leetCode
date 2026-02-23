package main

func main() {

}

func findAnagrams(s string, p string) []int {
    answer := []int{}
    ln := len(p)
    bp := []byte(p)
    slices.Sort(bp)
    for i := 0; i + (ln-1) < len(s); i++ {
        bs := []byte(s[i:i+ln])
        slices.Sort(bs)
        if bs[0] == bp[0] && eq(bs, bp) {
            answer = append(answer, i)
        }
    }
    return answer
}

func eq(s []byte, p []byte) bool{
    for i := range s {
        if s[i] != p[i] {
            return false
        }
    }
    return true
}
