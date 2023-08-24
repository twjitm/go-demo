package price

func getMaxPrice(prices []int64) int64 {
	var minIndex = 0
	var minValue = int64(0)
	var maxValue = int64(0)
	for index, price := range prices {
		if minValue == 0 || minValue > price {
			minValue = price
			minIndex = index
		}
		if minIndex < index && maxValue < price {
			maxValue = price
		}
	}
	return maxValue - minValue
}
