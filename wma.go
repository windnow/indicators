package indicators

func WMA(prices []float64, period int) []float64 {
	result := make([]float64, len(prices))
	if len(prices) < period {
		return result
	}

	denominator := float64(period * (period + 1) / 2)

	for i := period - 1; i < len(prices); i++ {
		numerator := 0.0
		for j := 0; j < period; j++ {
			weight := float64(j + 1)
			numerator += weight * prices[i-j]
		}
		result[i] = numerator / denominator
	}

	return result
}
