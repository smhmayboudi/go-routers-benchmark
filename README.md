# go-routers-benchmark

```shell
go test -benchmem -run=^$ -coverprofile=$(mktemp -t go-routers-benchmark) -bench . github.com/smhmayboudi/go-routers-benchmark
```

```shell
goos: darwin
goarch: arm64
pkg: github.com/smhmayboudi/go-routers-benchmark
BenchmarkKmux-8                  8289873               124.4 ns/op            64 B/op          3 allocs/op
BenchmarkKsmux-8                26064174                46.27 ns/op           18 B/op          1 allocs/op
BenchmarkChi-8                   7336424               151.0 ns/op           359 B/op          3 allocs/op
BenchmarkNetHTTP-8               6197923               193.8 ns/op            16 B/op          1 allocs/op
BenchmarkGin-8                  15284554                73.19 ns/op           65 B/op          1 allocs/op
BenchmarkEcho-8                 14089894                84.36 ns/op           27 B/op          1 allocs/op
BenchmarkKmuxWithParam-8         7409496               161.2 ns/op            84 B/op          3 allocs/op
BenchmarkKsmuxWithParam-8       16891990                73.16 ns/op           47 B/op          1 allocs/op
BenchmarkChiWithParam-8          5928310               192.1 ns/op           374 B/op          3 allocs/op
BenchmarkNetHTTPWithParam-8      1732419               688.1 ns/op           440 B/op         11 allocs/op
BenchmarkGinWithParam-8         10610310               105.5 ns/op            89 B/op          2 allocs/op
BenchmarkEchoWithParam-8        11384600               106.0 ns/op            55 B/op          2 allocs/op
BenchmarkKmuxWith2Param-8        6615656               181.1 ns/op            96 B/op          3 allocs/op
BenchmarkKsmuxWith2Param-8      13990076                89.06 ns/op           54 B/op          1 allocs/op
BenchmarkNetHTTPWith2Param-8     1368130               874.0 ns/op           536 B/op         13 allocs/op
BenchmarkGinWith2Param-8         9884214               123.0 ns/op           118 B/op          2 allocs/op
BenchmarkEchoWith2Param-8        9399381               129.1 ns/op            89 B/op          2 allocs/op
BenchmarkKmuxWith5Param-8        4949292               242.7 ns/op           134 B/op          3 allocs/op
BenchmarkKsmuxWith5Param-8       8175662               148.7 ns/op            97 B/op          1 allocs/op
BenchmarkNetHTTPWith5Param-8      852153              1374 ns/op             952 B/op         17 allocs/op
BenchmarkGinWith5Param-8         6331509               189.1 ns/op           164 B/op          2 allocs/op
BenchmarkEchoWith5Param-8        5907968               204.3 ns/op           154 B/op          2 allocs/op
PASS
coverage: 0.0% of statements
ok      github.com/smhmayboudi/go-routers-benchmark     31.154s
```