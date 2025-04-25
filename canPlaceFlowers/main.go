func canPlaceFlowers(flowerbed []int, n int) bool {
    if len(flowerbed) == 1 && flowerbed[0] == 0 {
        n--
    }
  
    if len(flowerbed) >= 2 && flowerbed[0] + flowerbed[1] == 0 {
        n--
    }
	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i]+flowerbed[i+1] == 0 {
			flowerbed[i+1] = 1
			n--
		}

		if flowerbed[i]+flowerbed[i+1] == 2 {
			flowerbed[i] = 0
			n++
		}
	}
	//	fmt.Println(n)
	return n <= 0
}
