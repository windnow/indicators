package indicators

import (
	"testing"

	"github.com/windnow/strategyapi/strategyapi"
)

func floatSliceAlmostEqual(a, b []float64, tolerance float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		diff := a[i] - b[i]
		if diff < 0 {
			diff = -diff
		}
		if diff > tolerance {
			return false
		}
	}
	return true
}

func TestEMA(t *testing.T) {
	prices := []float64{10, 11, 12, 13, 14, 15, 16}
	ema := EMA(prices, 3)
	if len(ema) != len(prices) {
		t.Errorf("EMA length mismatch")
	}
}

func TestSMA(t *testing.T) {
	prices := []float64{1, 2, 3, 4, 5}
	sma := SMA(prices, 3)
	expected := []float64{0, 0, 2, 3, 4}
	if !floatSliceAlmostEqual(sma, expected, 1e-6) {
		t.Errorf("SMA mismatch: got %v, expected %v", sma, expected)
	}
}

func TestRSI(t *testing.T) {
	prices := []float64{44, 44.15, 44.09, 44.15, 44.32, 44.96, 45.23, 45.12, 45.15, 45.33, 45.21, 45.10, 45.05, 44.94, 44.78}
	rsi := RSI(prices, 14)
	if len(rsi) != len(prices) {
		t.Errorf("RSI length mismatch")
	}
}

func TestMACD(t *testing.T) {
	prices := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	macd, signal, hist := MACD(prices, 3, 6, 2)
	if len(macd) != len(prices) || len(signal) != len(prices) || len(hist) != len(prices) {
		t.Errorf("MACD output length mismatch")
	}
}

func TestATR(t *testing.T) {
	candles := []strategyapi.Candle{
		{High: 12, Low: 10, Close: 11},
		{High: 13, Low: 10, Close: 12},
		{High: 14, Low: 11, Close: 13},
		{High: 15, Low: 12, Close: 14},
		{High: 16, Low: 13, Close: 15},
	}
	atr := ATR(candles, 3)
	if len(atr) != len(candles) {
		t.Errorf("ATR output length mismatch")
	}
}
