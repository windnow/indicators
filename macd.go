package indicators

func MACD(prices []float64, fast, slow, signal int) (macd, signalLine, hist []float64) {
	fastEMA := EMA(prices, fast)
	slowEMA := EMA(prices, slow)

	length := len(prices)
	macd = make([]float64, length)
	for i := 0; i < length; i++ {
		macd[i] = fastEMA[i] - slowEMA[i]
	}

	signalLine = EMA(macd, signal)

	hist = make([]float64, length)
	for i := 0; i < length; i++ {
		hist[i] = macd[i] - signalLine[i]
	}

	return
}
