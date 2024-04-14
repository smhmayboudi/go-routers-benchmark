# go-routers-benchmark

```shell
go test -benchmem -run=^$ -coverprofile=$(mktemp -t go-routers-benchmark) -bench . github.com/smhmayboudi/go-routers-benchmark
```

```shell
goos: darwin
goarch: arm64
pkg: github.com/smhmayboudi/go-routers-benchmark
BenchmarkKmux-8                	 8035486	       127.2 ns/op	      88 B/op	       3 allocs/op
BenchmarkKsmux-8               	26611471	        45.01 ns/op	      18 B/op	       1 allocs/op
BenchmarkChi-8                 	 7367400	       147.9 ns/op	     359 B/op	       3 allocs/op
BenchmarkNetHTTP-8             	 3582571	       311.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkGin-8                 	16156928	        72.47 ns/op	      64 B/op	       1 allocs/op
BenchmarkEcho-8                	12056065	        93.83 ns/op	      19 B/op	       1 allocs/op
BenchmarkFlow-8                	 5196439	       217.1 ns/op	     412 B/op	       2 allocs/op
BenchmarkFiber-8               	  806756	      1301 ns/op	    2106 B/op	      20 allocs/op
BenchmarkGorilla-8             	 2717814	       422.7 ns/op	     796 B/op	       7 allocs/op
BenchmarkKmuxWith1Param-8      	 5599077	       209.1 ns/op	     103 B/op	       5 allocs/op
BenchmarkKsmuxWith1Param-8     	 9791038	       120.3 ns/op	      75 B/op	       3 allocs/op
BenchmarkChiWith1Param-8       	 4883268	       238.6 ns/op	     411 B/op	       5 allocs/op
BenchmarkNetHTTPWith1Param-8   	 1798758	       669.4 ns/op	     440 B/op	      11 allocs/op
BenchmarkGinWith1Param-8       	 7796590	       154.1 ns/op	     114 B/op	       3 allocs/op
BenchmarkEchoWith1Param-8      	 8245624	       143.4 ns/op	      80 B/op	       3 allocs/op
BenchmarkFlowWith1Param-8      	 4462126	       261.8 ns/op	     478 B/op	       7 allocs/op
BenchmarkFiberWith1Param-8     	  891254	      1162 ns/op	    2097 B/op	      22 allocs/op
BenchmarkGorillaWith1Param-8   	 1965139	       586.1 ns/op	    1138 B/op	       9 allocs/op
BenchmarkKmuxWith2Param-8      	 4524276	       262.8 ns/op	     179 B/op	       6 allocs/op
BenchmarkKsmuxWith2Param-8     	 6810192	       173.2 ns/op	     119 B/op	       4 allocs/op
BenchmarkChiWith2Param-8       	 3791394	       309.2 ns/op	     486 B/op	       6 allocs/op
BenchmarkNetHTTPWith2Param-8   	 1354172	       886.9 ns/op	     536 B/op	      13 allocs/op
BenchmarkGinWith2Param-8       	 5612060	       211.5 ns/op	     159 B/op	       4 allocs/op
BenchmarkEchoWith2Param-8      	 5934452	       201.4 ns/op	     125 B/op	       4 allocs/op
BenchmarkFlowWith2Param-8      	 2821000	       418.9 ns/op	     623 B/op	      12 allocs/op
BenchmarkFiberWith2Param-8     	  753645	      1374 ns/op	    2210 B/op	      26 allocs/op
BenchmarkGorillaWith2Param-8   	 1462051	       812.4 ns/op	    1182 B/op	      10 allocs/op
BenchmarkKmuxWith5Param-8      	 2738371	       436.6 ns/op	     338 B/op	       9 allocs/op
BenchmarkKsmuxWith5Param-8     	 3546912	       336.6 ns/op	     251 B/op	       7 allocs/op
BenchmarkChiWith5Param-8       	 2134111	       560.9 ns/op	     637 B/op	       9 allocs/op
BenchmarkNetHTTPWith5Param-8   	  782593	      1473 ns/op	    1016 B/op	      17 allocs/op
BenchmarkGinWith5Param-8       	 3078026	       391.5 ns/op	     295 B/op	       7 allocs/op
BenchmarkEchoWith5Param-8      	 3111195	       388.0 ns/op	     262 B/op	       7 allocs/op
BenchmarkFlowWith5Param-8      	 1222584	       936.4 ns/op	    1069 B/op	      27 allocs/op
BenchmarkFiberWith5Param-8     	  740115	      1390 ns/op	    2212 B/op	      26 allocs/op
BenchmarkGorillaWith5Param-8   	  700424	      1564 ns/op	    1328 B/op	      13 allocs/op
PASS
coverage: 0.0% of statements
ok  	github.com/smhmayboudi/go-routers-benchmark	52.850s
```