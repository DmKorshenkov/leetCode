func stringHash(s string, k int) string {
	bb := bytes.Buffer{}
	bb2 := bytes.Buffer{}
	bb.WriteString(s)

	for i := 0; i < len(s); i += k {
		tmp := helper(bb.Bytes()[i : i+k])
		//		fmt.Println(tmp)
		bb2.WriteByte(tmp + 97)

	}

	return bb2.String()
}

func helper(arr []byte) byte {
	var res int
//	var check byte
	res = 0

	for i := range arr {
		//		fmt.Println("s = ", string(arr[i]), "ascii = ", arr[i]-97)
		res += int(arr[i] - 97)
	}
	return byte(res % 26)
}
