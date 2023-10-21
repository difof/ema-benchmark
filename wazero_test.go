package main

import (
	"context"
	"testing"
)

func TestEMAWasm(t *testing.T) {
	prices := GeneratePrices(100_000_000)
	period := 10

	mod, err := LoadWasmModule()
	if err != nil {
		t.Fatalf("failed to load wasm module: %v", err)
	}

	result := EMAWasm(context.Background(), mod, prices, period)
	if len(result) != len(prices) {
		t.Errorf("expected %d results, got %d", len(prices), len(result))
	}

	//expected := EMAGo(prices, period)

	//t.Logf("%v", prices)
	//t.Logf("%v", result)
	//t.Logf("%v", expected)
}
