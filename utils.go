package indicators

func Average(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}

	return sum / float64(len(values))
}

func maxf(a, b, c float64) float64 {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}

func absf(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}
