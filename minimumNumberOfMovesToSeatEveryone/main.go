func minMovesToSeat(seats []int, students []int) int {
    n := len(seats)
    // sort seats
    for i := 1; i < n; i++ {
        key := seats[i]

        j := i-1
        for j >= 0 && seats[j] > key {
            seats[j+1] = seats[j] 
            j--
        }
        seats[j+1] = key
    }
    // sort student 
    for i := 1; i < n; i++ {
        key := students[i]

        j := i-1
        for j >= 0 && students[j] > key {
            students[j+1] = students[j]
            j--
        }
        students[j+1] = key
    }
    res := 0
    for n > 0 {
        count := students[n-1] - seats[n-1]
        abc(&count)
        res += count 
        //res += max(students[n-1], seats[n-1]) - min(students[n-1], seats[n-1])
        n--
    }
 //   fmt.Println(seats, students)
    return res
}

func abc(n *int) {
    num := *n
    if num < 0 {
        num = -num
    }
    *n = num
}
