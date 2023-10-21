package main

import (
	"context"
	"github.com/difof/goul/errors"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"log"
	"os"
	"unsafe"
)

func LoadWasmModule() (mod api.Module, err error) {
	filename := "assemblyscript/build/release.wasm"

	var modBytes []byte
	modBytes, err = os.ReadFile(filename)
	if err != nil {
		err = errors.Newif(err, "failed to read file: %s", filename)
		return
	}

	ctx := context.Background()

	rt := wazero.NewRuntimeWithConfig(
		ctx,
		wazero.NewRuntimeConfigCompiler(),
	)

	_, err = rt.NewHostModuleBuilder("env").Instantiate(ctx)
	if err != nil {
		err = errors.Newif(err, "failed to instantiate host module: %s", filename)
		return
	}

	mod, err = rt.InstantiateWithConfig(
		ctx,
		modBytes,
		wazero.NewModuleConfig(),
	)
	if err != nil {
		err = errors.Newif(err, "failed to instantiate module: %s", filename)
		return
	}

	return
}

func EMAWasm(ctx context.Context, mod api.Module, prices []float64, period int) []float64 {
	mem := mod.ExportedMemory("memory")
	pricesBytes := float64ToBytesLE(prices)

	if mem.Size() < uint32(len(pricesBytes)*2) {
		mem.Grow(szToPageSize(len(pricesBytes) * 2))
	}
	mem.Write(0, pricesBytes)

	//start := time.Now()
	emaFunc := mod.ExportedFunction("ema")
	results, err := emaFunc.Call(ctx, 0, uint64(len(pricesBytes)), uint64(period))
	if err != nil {
		log.Fatalf("failed to call main function: %s", err)
		return nil
	}
	//fmt.Printf("[%dus] AOT WASM EMA (function call only)\n", time.Since(start).Microseconds())

	if len(results) != 1 {
		log.Fatalf("unexpected number of results: %d", len(results))
		return nil
	}

	resultPtr := results[0]
	resultBytes, ok := mem.Read(uint32(resultPtr), uint32(len(pricesBytes)))
	if !ok {
		log.Fatalf("failed to read memory")
		return nil
	}

	return bytesToFloat64LE(resultBytes)
}

// szToPageSize converts the given size to page size (64kb)
func szToPageSize(sz int) uint32 {
	return uint32((sz + 65535) / 65536)
}

func float64ToBytesLE(f []float64) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&f[0])), len(f)*8)

	//b := make([]byte, len(f)*8)
	//for i, v := range f {
	//	binary.LittleEndian.PutUint64(b[i*8:], math.Float64bits(v))
	//}
	//
	//return b
}

func bytesToFloat64LE(b []byte) []float64 {
	return unsafe.Slice((*float64)(unsafe.Pointer(&b[0])), len(b)/8)

	//f := make([]float64, len(b)/8)
	//for i := 0; i < len(f); i++ {
	//	f[i] = math.Float64frombits(binary.LittleEndian.Uint64(b[i*8 : (i+1)*8]))
	//}
	//
	//return f
}
