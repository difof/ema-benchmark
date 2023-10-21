# EMA Benchmark

I was interested by the [wazero](https://github.com/tetratelabs/wazero) project and I wanted to
benchmark the performance by calculating EMA of random data.

EMA is calculated in both native Go and WASM.
Performance result of WASM is almost same as native Go with
around 100 microseconds difference.

The only bottlenecks are wazero's module memory operations and
user data encoding/decoding. So I sacrificed portability and
a bit of safety to speed up encoding.

## Results

Machine MBP M1 2020 16GB, EMA of 100,000 data points:

| | | | | |
|--------------------|----|-----------------|----------------|--------------|
| BenchmarkEMAGo-8   | 32 | 364697573 ns/op | 800006150 B/op | 1 allocs/op  |
| BenchmarkEMAWasm-8 | 20 | 504017300 ns/op | 80010101 B/op  | 12 allocs/op |

Running `main.go`:

```
[1106us] Generating 100000 price points
[426us] Native Go EMA
[685us] AOT WASM EMA (call + memory encoding/decoding)
```

## How to run

#### Dependencies:

- Go
- Node
- NPM
- Git

```bash
git clone https://github.com/difof/wasm_ema_benchmark
cd wasm_ema_benchmark/assemblyscript
npm install
npm run asbuild
cd ../go
go run main.go
go test -bench=. -benchmem -benchtime=10s
```

# License
MIT