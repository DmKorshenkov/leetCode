func minPartitions(n string) int {
    var a int
    for _, r := range n {
        if int(r) - 48 > a {
            a = int(r) - 48
        }
    }  
    return a   
}
