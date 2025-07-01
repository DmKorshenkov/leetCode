func isPerfectSquare(num int) bool {
    for i := 1; i <= num -(num/2) && i * i <= num;i++{
        if i * i == num {
            return true
        }
    }  
    return false
}
