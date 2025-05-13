func asteroidsDestroyed(mass int, asteroids []int) bool {
    sort.Ints(asteroids)

    for i := range asteroids {
        if asteroids[i] > mass {
            return false
        }
        mass += asteroids[i]
    }

    return true
}
