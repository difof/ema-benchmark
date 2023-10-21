package main

import (
	"context"
	"testing"
)

var (
	prices = GeneratePrices(100_000_000)
	period = 10
)

func BenchmarkEMAGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EMAGo(prices, period)
	}
}

func BenchmarkEMAWasm(b *testing.B) {
	mod, err := LoadWasmModule()
	if err != nil {
		b.Fatalf("failed to load wasm module: %v", err)
	}

	for i := 0; i < b.N; i++ {
		EMAWasm(context.Background(), mod, prices, period)
	}
}
