func destCity(paths [][]string) string {
    mp := make(map[string]bool)

    for i := range paths {
        if _, ok := mp[paths[i][0]]; ok {
            mp[paths[i][0]] = true
        } else {
            mp[paths[i][0]] = true
        }
        if _, ok := mp[paths[i][1]]; !ok {
            mp[paths[i][1]] = false
        }
    }
    for key, _ :=range mp {
        if mp[key] == false {
            return key
        }
    }
    return ""
}
