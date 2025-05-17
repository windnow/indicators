package indicators

func SMA(prices []float64, period int) []float64 {
	result := make([]float64, len(prices))
	if len(prices) < period {
		return result
	}

	sum := 0.0
	for i := 0; i < period; i++ {
		sum += prices[i]
	}

	result[period-1] = sum / float64(period)

	for i := period; i < len(prices); i++ {
		sum = sum - prices[i-period] + prices[i]
		result[i] = sum / float64(period)
	}

	return result
}
