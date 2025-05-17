package indicators

func RSI(prices []float64, period int) []float64 {
	rsis := make([]float64, len(prices))
	if len(prices) < period+1 {
		return rsis
	}

	var gain, loss float64
	for i := 1; i <= period; i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			gain += diff
		} else {
			loss -= diff
		}
	}

	avgGain := gain / float64(period)
	avgLoss := loss / float64(period)
	rsis[period] = 100 - (100 / (1 + avgGain/avgLoss))

	for i := period + 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			avgGain = (avgGain*(float64(period-1)) + diff) / float64(period)
			avgLoss = (avgLoss * float64(period-1)) / float64(period)
		} else {
			avgGain = (avgGain * float64(period-1)) / float64(period)
			avgLoss = (avgLoss*(float64(period-1)) - diff) / float64(period)
		}
		rsis[i] = 100 - (100 / (1 + avgGain/avgLoss))
	}

	return rsis
}
