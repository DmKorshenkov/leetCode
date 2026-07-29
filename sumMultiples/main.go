func sumOfMultiples(n int) int {
    help := func(k int) int{
        return k * (n/k)*((n/k)+1) / 2
    }

    return help(3)+help(5)+help(7)-help(15)-help(21)-help(35)+help(105)
}
