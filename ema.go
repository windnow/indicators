package indicators

func EMA(prices []float64, period int) []float64 {
	result := make([]float64, len(prices))
	if len(prices) < period {
		return result
	}

	alpha := 2.0 / float64(period+1)
	result[period-1] = Average(prices[:period])

	for i := period; i < len(prices); i++ {
		result[i] = alpha*prices[i] + (1-alpha)*result[i-1]
	}

	return result
}