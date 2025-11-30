
func maxProfit(prices []int) int {
	maxProfit := 0
 //   win := 0
    price := prices[0]
    for i := 1; i < len(prices); i++ {
        if prices[i] < price {
            price = prices[i]
        }
        if prices[i] - price > maxProfit {
            maxProfit = prices[i] - price
        }
    }
    return maxProfit
}
