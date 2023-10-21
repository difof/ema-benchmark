package main

import (
	"math/rand"
	"time"
)

func GeneratePrices(n int) []float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	prices := make([]float64, n)

	initialPrice := 30000.0
	volatility := .01
	stepSize := .1

	price := initialPrice
	for i := 0; i < n; i++ {
		price += stepSize * r.NormFloat64()
		price *= 1 + (volatility * r.NormFloat64())

		prices[i] = price
	}

	return prices
}
