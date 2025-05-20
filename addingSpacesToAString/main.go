func addSpaces(s string, spaces []int) string {    
    bb := bytes.Buffer{}
    
    inS := 0
    inSpaces := 0
    for inS < len(s){
        if inSpaces < len(spaces) && inS == spaces[inSpaces]{
            bb.WriteByte(byte(' '))
            inSpaces++
        }
        bb.WriteByte(s[inS])
        inS++
    }

    if inSpaces < len(spaces) {
        for ;inSpaces < len(spaces); bb.WriteByte(byte(' ')) {
            inSpaces++
        }
    }
    return bb.String()
}
