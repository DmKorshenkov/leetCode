func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }

	arr := make([]int, len(num1)+len(num2), len(num1)+len(num2))
	num1 = reversString(num1)
	num2 = reversString(num2)
    
	for in1 := 0; in1 < len(num1); in1++ {

		for in2 := 0; in2 < len(num2); in2++ {
			m := int((num1[in1] - 48) * (num2[in2] - 48))
			arr[in2+in1] += m
			if arr[in2+in1] >= 10 {
				arr[in2+in1+1] += arr[in2+in1] / 10
				arr[in2+in1] = arr[in2+in1] % 10
			}
		}
	}
	bb := bytes.Buffer{}
	for i := 0; i < len(arr)-1; i++ {
		bb.WriteByte(byte(arr[i] + 48))
	}
	if arr[len(arr)-1] != 0 {
		bb.WriteByte(byte(arr[len(arr)-1] + 48))
	}
	res := reversString(bb.String())
	return res
}

func reversString(str string) string {
	if len(str) <= 1 {
		return str
	}
	return reversString(str[1:]) + string(str[0])
}
