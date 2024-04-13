# go-routers-benchmark

```shell
go test -benchmem -run=^$ -coverprofile=$(mktemp -t go-routers-benchmark) -bench . github.com/smhmayboudi/go-routers-benchmark
```

```shell
goos: darwin
goarch: arm64
pkg: github.com/smhmayboudi/go-routers-benchmark
BenchmarkKmux-8                	 7788574	       129.3 ns/op	      89 B/op	       3 allocs/op
BenchmarkKsmux-8               	26224276	        44.79 ns/op	      18 B/op	       1 allocs/op
BenchmarkChi-8                 	 7368996	       149.9 ns/op	     359 B/op	       3 allocs/op
BenchmarkNetHTTP-8             	 3888301	       308.3 ns/op	      22 B/op	       1 allocs/op
BenchmarkGin-8                 	16242682	        72.08 ns/op	      64 B/op	       1 allocs/op
BenchmarkEcho-8                	12776564	        93.34 ns/op	      18 B/op	       1 allocs/op
BenchmarkFlow-8                	 5160415	       218.0 ns/op	     413 B/op	       2 allocs/op
BenchmarkKmuxWith1Param-8      	 5628194	       210.1 ns/op	     103 B/op	       5 allocs/op
BenchmarkKsmuxWith1Param-8     	 9794695	       120.8 ns/op	      75 B/op	       3 allocs/op
BenchmarkChiWith1Param-8       	 4846485	       237.4 ns/op	     411 B/op	       5 allocs/op
BenchmarkNetHTTPWith1Param-8   	 1790638	       671.1 ns/op	     440 B/op	      11 allocs/op
BenchmarkGinWith1Param-8       	 7960473	       150.8 ns/op	     113 B/op	       3 allocs/op
BenchmarkEchoWith1Param-8      	 8274974	       144.4 ns/op	      80 B/op	       3 allocs/op
BenchmarkFlowWith1Param-8      	 4459946	       260.9 ns/op	     478 B/op	       7 allocs/op
BenchmarkKmuxWith2Param-8      	 4546233	       436.1 ns/op	     179 B/op	       6 allocs/op
BenchmarkKsmuxWith2Param-8     	 6234756	       171.7 ns/op	     123 B/op	       4 allocs/op
BenchmarkChiWith2Param-8       	 3861104	       305.2 ns/op	     485 B/op	       6 allocs/op
BenchmarkNetHTTPWith2Param-8   	 1350007	       886.8 ns/op	     536 B/op	      13 allocs/op
BenchmarkGinWith2Param-8       	 5889428	       206.4 ns/op	     157 B/op	       4 allocs/op
BenchmarkEchoWith2Param-8      	 5967622	       200.6 ns/op	     124 B/op	       4 allocs/op
BenchmarkFlowWith2Param-8      	 5393534	       222.6 ns/op	     161 B/op	       4 allocs/op
BenchmarkKmuxWith5Param-8      	 2757657	       433.7 ns/op	     337 B/op	       9 allocs/op
BenchmarkKsmuxWith5Param-8     	 3523689	       334.8 ns/op	     252 B/op	       7 allocs/op
BenchmarkChiWith5Param-8       	 2132821	       563.8 ns/op	     637 B/op	       9 allocs/op
BenchmarkNetHTTPWith5Param-8   	  791149	      1457 ns/op	    1016 B/op	      17 allocs/op
BenchmarkGinWith5Param-8       	 3098656	       390.7 ns/op	     294 B/op	       7 allocs/op
BenchmarkEchoWith5Param-8      	 3132655	       380.0 ns/op	     261 B/op	       7 allocs/op
BenchmarkFlowWith5Param-8      	 1318240	       905.4 ns/op	    1061 B/op	      27 allocs/op
PASS
coverage: 0.0% of statements
ok  	github.com/smhmayboudi/go-routers-benchmark	43.722s
```