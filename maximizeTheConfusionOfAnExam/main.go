func maxConsecutiveAnswers(answerKey string, k int) int {
    var res int
    var t, f int //window
    var l int

    for r := 0; r < len(answerKey); r++ {
        if rune(answerKey[r]) == 'T'{
            t++
        } else {
            f++
        }

        for min(t,f) > k {        
            if rune(answerKey[l]) == 'T' {
                t--
            } else {
                f--
            }
            l++
        }

        if t+f > res {
            res = t+f
        }
    }
//	fmt.Println(res)
	return res
}
