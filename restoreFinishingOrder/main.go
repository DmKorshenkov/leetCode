func recoverOrder(order []int, friends []int) []int {
    for i, step := 0, 0; i < len(order); i++ {
        for j := 0; j < len(friends); j++ {
            if order[i] == friends[j] {
                friends[step], friends[j] = friends[j],friends[step]
                step++
            }
        }
    }
    return friends
}
