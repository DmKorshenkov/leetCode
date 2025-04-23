func reverseVowels(s string) string {
  var i, j = 0, len(s)-1
  bb := bytes.Buffer{}
  bb.WriteString(s)
  
  for i <= j {


    if !check(bb.Bytes()[i]) {
        i++
        continue
    }
    if !check(bb.Bytes()[j]) {
        j--
        continue
    }

    bb.Bytes()[i],bb.Bytes()[j] = bb.Bytes()[j],bb.Bytes()[i]
    i++
    j--
  }
  return bb.String()  
}

func check(n byte) bool {
  check := "aeiouAEIOU"
  for i := range check {
    if n == check[i] {
        return true
    }
  }

  return false  

