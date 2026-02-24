package main

func main() {

}
func findAnagrams(s string, p string) []int {
    answer := []int{}
    bp := make([]int, 26)
    bs := make([]int, 26)
    lnP := len(p)
    for i := range p {
        bp[p[i]-97]++
    }

    for i := 0; i < len(s); i++ {
        c := int(s[i]-97)
        bs[c]++
        if i >= lnP {
            bs[int(s[i-lnP]-97)]--
        }
        if bp[c] > 0 && eq(bs, bp) {
            answer = append(answer, i-lnP+1)
        }
    }
    return answer
}

func eq(s []int, p []int) bool{
    for i := range s {
        if s[i] != p[i] {
            return false
        }
    }
    return true
}
