package main

func main(){

}

func checkOnesSegment(s string) bool {
    i := 0

    for i < len(s) && s[i] == '1' {
        i++
    }

    for i < len(s) {
        if s[i] == '1' {
            return false
        }
        i++
    }
    return true
}
