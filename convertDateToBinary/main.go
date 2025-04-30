func convertDateToBinary(date string) string {
	dateSl := strings.Split(date, "-")
	bb := bytes.Buffer{}

	for _, str := range dateSl {
		num, _ := strconv.Atoi(str)
		binNum:= fmt.Sprintf("%b", num)
		bb.WriteString(binNum + "-")
	}
	return (string(bb.Bytes()[:bb.Len()-1]))
//	return ""
}
