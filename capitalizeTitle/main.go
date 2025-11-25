func capitalizeTitle(title string) string {
	bb := bytes.Buffer{}
	res := bytes.Buffer{}
	for i := 0; i < len(title); i++ {
		bb.WriteByte(title[i])
		if (title[i] == byte(' ')) || i == len(title)-1 {
			tmp := trimSpace(bb.Bytes())

			toLowerCaseHelp(tmp)
			if len(tmp) > 2 {
				tmp[0] -= 32
			}
			//res.WriteString(string(tmp))
			//fmt.Println(res.String())
			//fmt.Println(string(tmp))
			res.WriteString(string(tmp))
			res.WriteByte(byte(' '))
			bb.Reset()
		}

	}

	//fmt.Println(res.String())
	return string(res.Bytes())
}

func toLowerCaseHelp(s []byte) string {
	var res string
	for i := 0; i < len(s); i++ {
		if s[i] >= byte('A') && s[i] <= byte('Z') {
			s[i] += 32
		}
	}
	res = string(s)
	return res
}

func trimSpace(s []byte) []byte {

	res := bytes.Buffer{}
	for i := 0; i < len(s); i++ {
		if s[i] != byte(' ') {
			res.WriteByte(s[i])
		}
	}

	//	fmt.Println(string(s))
	return res.Bytes()
}
