func angleCLock(hour int, minutes int) float64 {
	return abc(30.0*float64(hour)-5.5*float64(minutes))
}

func abc(n float64) float64 {
	if n < 0 {
		n = (-n)
	}
	if n > 180 {
		n = 360-n
	}
	return n
}
