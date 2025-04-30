func interpret(command string) string {
  bb := bytes.Buffer{}  

  for i := 0; i < len(command); i++ {
    if command[i] == byte('G') {
            bb.WriteRune('G')
            continue 
    }
    
    if command[i] == byte('(') && command[i+1] == byte('a') {
            bb.WriteString("al")
            i += 3
    }
    if command[i] == byte('(') && command[i+1] == byte(')') {
            bb.WriteRune('o')
            i += 1
    }
    
    }
    return bb.String()
} 
