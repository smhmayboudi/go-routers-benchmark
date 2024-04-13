# go-routers-benchmark

```shell
go test -benchmem -run=^$ -coverprofile=$(mktemp -t go-routers-benchmark) -bench . github.com/smhmayboudi/go-routers-benchmark
```

```shell
goos: darwin
goarch: arm64
pkg: github.com/smhmayboudi/go-routers-benchmark
BenchmarkKmux-8                	 9280707	       127.5 ns/op	      86 B/op	       3 allocs/op
BenchmarkKsmux-8               	26013181	        45.35 ns/op	      18 B/op	       1 allocs/op
BenchmarkChi-8                 	 7314086	       159.2 ns/op	     359 B/op	       3 allocs/op
BenchmarkNetHTTP-8             	 3862918	       310.2 ns/op	      22 B/op	       1 allocs/op
BenchmarkGin-8                 	16447662	        72.77 ns/op	      64 B/op	       1 allocs/op
BenchmarkEcho-8                	12708246	        92.67 ns/op	      18 B/op	       1 allocs/op
BenchmarkFlow-8                	 5172818	       216.0 ns/op	     413 B/op	       2 allocs/op
BenchmarkFiber-8               	  791839	      1329 ns/op	    2108 B/op	      20 allocs/op
BenchmarkKmuxWith1Param-8      	 5525716	       212.1 ns/op	     104 B/op	       5 allocs/op
BenchmarkKsmuxWith1Param-8     	 9695960	       122.5 ns/op	      75 B/op	       3 allocs/op
BenchmarkChiWith1Param-8       	 4771207	       240.5 ns/op	     412 B/op	       5 allocs/op
BenchmarkNetHTTPWith1Param-8   	 1786905	       672.9 ns/op	     440 B/op	      11 allocs/op
BenchmarkGinWith1Param-8       	 7865022	       150.3 ns/op	     114 B/op	       3 allocs/op
BenchmarkEchoWith1Param-8      	 8255751	       143.8 ns/op	      80 B/op	       3 allocs/op
BenchmarkFlowWith1Param-8      	 4471314	       261.3 ns/op	     478 B/op	       7 allocs/op
BenchmarkFiberWith1Param-8     	 1041735	      1143 ns/op	    2077 B/op	      22 allocs/op
BenchmarkKmuxWith2Param-8      	 4497944	       264.3 ns/op	     179 B/op	       6 allocs/op
BenchmarkKsmuxWith2Param-8     	 6803762	       172.6 ns/op	     119 B/op	       4 allocs/op
BenchmarkChiWith2Param-8       	 3710054	       311.6 ns/op	     452 B/op	       6 allocs/op
BenchmarkNetHTTPWith2Param-8   	 1352978	       886.5 ns/op	     536 B/op	      13 allocs/op
BenchmarkGinWith2Param-8       	 5649060	       213.3 ns/op	     159 B/op	       4 allocs/op
BenchmarkEchoWith2Param-8      	 5996229	       200.1 ns/op	     124 B/op	       4 allocs/op
BenchmarkFlowWith2Param-8      	 2799889	       421.3 ns/op	     623 B/op	      12 allocs/op
BenchmarkFiberWith2Param-8     	  723422	      1387 ns/op	    2216 B/op	      26 allocs/op
BenchmarkKmuxWith5Param-8      	 2731929	       437.7 ns/op	     338 B/op	       9 allocs/op
BenchmarkKsmuxWith5Param-8     	 3527212	       338.3 ns/op	     252 B/op	       7 allocs/op
BenchmarkChiWith5Param-8       	 2117115	       563.6 ns/op	     638 B/op	       9 allocs/op
BenchmarkNetHTTPWith5Param-8   	  809256	      1463 ns/op	    1016 B/op	      17 allocs/op
BenchmarkGinWith5Param-8       	 3139483	       383.2 ns/op	     293 B/op	       7 allocs/op
BenchmarkEchoWith5Param-8      	 3143030	       378.3 ns/op	     261 B/op	       7 allocs/op
BenchmarkFlowWith5Param-8      	 1309200	       912.9 ns/op	    1062 B/op	      27 allocs/op
BenchmarkFiberWith5Param-8     	  741900	      1395 ns/op	    2212 B/op	      26 allocs/op
PASS
coverage: 0.0% of statements
ok  	github.com/smhmayboudi/go-routers-benchmark	50.418s
```