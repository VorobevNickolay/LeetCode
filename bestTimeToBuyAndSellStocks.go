package LeetCode

func maxProfit(prices []int) int {
	prof := 0
	b := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < b {
			b = prices[i]
		} else {
			if prof < prices[i]-b {
				prof = prices[i] - b
			}
		}
	}
	return prof

}
