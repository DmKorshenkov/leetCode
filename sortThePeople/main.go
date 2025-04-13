func sortPeople(names []string, heights []int) []string {

    for i := 0; i < len(names); i++{

        for j := 0;  j < len(names); j++ {
            if heights[i] > heights[j] {
                names[i],names[j] = names[j],names[i]
                heights[i],heights[j] = heights[j],heights[i]
                
            }
        }
    }
    return names
}
