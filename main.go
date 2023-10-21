package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	numPrices := 100_000
	period := 10

	var prices []float64
	fmt.Printf("[%dus] Generating %d price points\n",
		MeasureTime(func() {
			prices = GeneratePrices(numPrices)
		}).Microseconds(),
		numPrices,
	)

	fmt.Printf("[%dus] Native Go EMA\n",
		MeasureTime(func() {
			EMAGo(prices, period)
		}).Microseconds(),
	)

	mod, err := LoadWasmModule()
	if err != nil {
		panic(err)
	}

	fmt.Printf("[%dus] AOT WASM EMA (call + memory encoding/decoding)\n",
		MeasureTime(func() {
			EMAWasm(context.Background(), mod, prices, period)
		}).Microseconds(),
	)
}

func MeasureTime(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Since(start)
}
