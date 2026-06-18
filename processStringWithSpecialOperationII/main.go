func processStr(s string, k int64) byte {
	var size int64

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '*':
			if size > 0 {
				size -= 1
			}
		case '#':
			size *= 2
		case '%':
		default:
			size++
		}
	}
	if k >= size {
		return '.'
	}
	for i := len(s) - 1; ; i-- {
		switch s[i] {
		case '*':
			size += 1
		case '#':
			size /= 2
			k %= size
		case '%':
			k = size - 1 - k
		default:
			size--
			if k == size {
				return s[i]
			}
		}
	}
}
