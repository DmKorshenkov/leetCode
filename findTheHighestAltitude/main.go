func largestAltitude(gain []int) int {
    var altitude, res int

    altitude = 0
    res = 0

    for i := range gain {
     
        altitude += gain[i]
     
        if altitude > res {
            res = altitude
        }

    }
    return res
}
