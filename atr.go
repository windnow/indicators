package indicators

import "github.com/windnow/strategyapi/strategyapi"

func ATR(candles []strategyapi.Candle, period int) []float64 {
	trs := make([]float64, len(candles))
	for i := 1; i < len(candles); i++ {
		high := candles[i].High
		low := candles[i].Low
		prevClose := candles[i-1].Close

		tr := maxf(
			high-low,
			absf(high-prevClose),
			absf(low-prevClose),
		)
		trs[i] = tr
	}

	return EMA(trs, period)
}
