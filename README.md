# copy-bench

Benchmarks for [Copy](https://github.com/gotidy/copy) project

```sh
go test -bench=. -benchmem ./...
goos: darwin
goarch: amd64
pkg: github.com/gotidy/copy-bench
BenchmarkManualCopy-12         177310519                6.92 ns/op             0 B/op          0 allocs/op
BenchmarkCopiers-12             13476417                84.1 ns/op             0 B/op          0 allocs/op
BenchmarkCopier-12              40226689                27.5 ns/op             0 B/op          0 allocs/op
BenchmarkJinzhuCopier-12          407480                2711 ns/op          2480 B/op         34 allocs/op
BenchmarkDeepcopier-12            262836                4346 ns/op          4032 B/op         73 allocs/op
PASS
ok      github.com/gotidy/copy-bench    6.922s
```
