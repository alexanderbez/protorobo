# protorobo

A quick n dirty benchmark comparing serialization between Protobuf's and
Ethereum's RLP protocols.

Serialization is done against a simple type, `MyMessage`, containing a few byte
slice fields representing different properties that could be common in a
blockchain implementation.

```go
type MyMessage struct {
	Address []byte 
	Hash    []byte 
	Code    []byte
}
```

## Benchmarks

```shell
$ make bench
```

```shell
goos: darwin
goarch: amd64
pkg: github.com/alexanderbez/protorobo/types
PASS
benchmark                          iter       time/iter
---------                          ----       ---------
BenchmarkGogoProtoTinyEncode-4    10000    113.35 μs/op
BenchmarkGogoProtoLargeEncode-4    3000    478.34 μs/op
BenchmarkGogoProtoTinyDecode-4    20000     82.79 μs/op
BenchmarkGogoProtoLargeDecode-4    5000    213.69 μs/op
BenchmarkETHRLPTinyEncode-4       10000    165.47 μs/op
BenchmarkETHRLPLargeEncode-4       2000    697.68 μs/op
BenchmarkETHRLPTinyDecode-4       10000    121.46 μs/op
BenchmarkETHRLPLargeDecode-4       3000    500.55 μs/op
ok      github.com/alexanderbez/protorobo/types 12.158s
```

__Note__: Serialization in all benchmarks is done against a pre-instantiated
type. In other words, each benchmark run is using the same data. Perhaps it'd be
wise to modify the benchmarking to create a new type with randomly seeded data
for reach run excluding the time it takes to create/seed said type.
