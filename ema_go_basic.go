package main

func EMAGo(data []float64, period int) []float64 {
	k := 2.0 / float64(period+1)
	ema := make([]float64, len(data))
	ema[0] = data[0]

	for i := 1; i < len(data); i++ {
		ema[i] = (data[i] * k) + (ema[i-1] * (1 - k))
	}

	return ema
}
