# go-routers-benchmark

```shell
go test -benchmem -run=^$ -coverprofile=$(mktemp -t go-routers-benchmark) -bench . github.com/smhmayboudi/go-routers-benchmark
```

```shell
goos: darwin
goarch: arm64
pkg: github.com/smhmayboudi/go-routers-benchmark
BenchmarkKmux-8                  	 8005923	       128.1 ns/op	      88 B/op	       3 allocs/op
BenchmarkKsmux-8                 	26106228	        45.39 ns/op	      18 B/op	       1 allocs/op
BenchmarkChi-8                   	 7405846	       147.4 ns/op	     359 B/op	       3 allocs/op
BenchmarkNetHTTP-8               	 3892263	       310.1 ns/op	      22 B/op	       1 allocs/op
BenchmarkGin-8                   	16070746	        72.36 ns/op	      64 B/op	       1 allocs/op
BenchmarkEcho-8                  	12634437	        93.60 ns/op	      18 B/op	       1 allocs/op
BenchmarkFlow-8                  	 5216252	       215.0 ns/op	     412 B/op	       2 allocs/op
BenchmarkFiber-8                 	  792627	      1315 ns/op	    2108 B/op	      20 allocs/op
BenchmarkGorilla-8               	 2735756	       421.0 ns/op	     796 B/op	       7 allocs/op
BenchmarkBunRouter-8             	30482545	        39.72 ns/op	      17 B/op	       0 allocs/op

BenchmarkKmuxWith1Param-8        	 5757566	       209.0 ns/op	     103 B/op	       5 allocs/op
BenchmarkKsmuxWith1Param-8       	 9776998	       120.9 ns/op	      75 B/op	       3 allocs/op
BenchmarkChiWith1Param-8         	 4805173	       238.0 ns/op	     411 B/op	       5 allocs/op
BenchmarkNetHTTPWith1Param-8     	 1799743	       666.6 ns/op	     440 B/op	      11 allocs/op
BenchmarkGinWith1Param-8         	 7541032	       156.7 ns/op	     115 B/op	       3 allocs/op
BenchmarkEchoWith1Param-8        	 8225959	       144.1 ns/op	      80 B/op	       3 allocs/op
BenchmarkFlowWith1Param-8        	 4476498	       261.0 ns/op	     478 B/op	       7 allocs/op
BenchmarkFiberWith1Param-8       	  921835	      1146 ns/op	    2092 B/op	      22 allocs/op
BenchmarkGorillaWith1Param-8     	 2022306	       583.5 ns/op	    1137 B/op	       9 allocs/op
BenchmarkBunRouterWith1Param-8   	 5575536	       211.8 ns/op	     112 B/op	       4 allocs/op

BenchmarkKmuxWith2Param-8        	 4541538	       262.3 ns/op	     179 B/op	       6 allocs/op
BenchmarkKsmuxWith2Param-8       	 6868298	       171.5 ns/op	     119 B/op	       4 allocs/op
BenchmarkChiWith2Param-8         	 3798561	       316.9 ns/op	     486 B/op	       6 allocs/op
BenchmarkNetHTTPWith2Param-8     	 1359296	       889.2 ns/op	     536 B/op	      13 allocs/op
BenchmarkGinWith2Param-8         	 5548041	       213.3 ns/op	     160 B/op	       4 allocs/op
BenchmarkEchoWith2Param-8        	 5878262	       203.3 ns/op	     125 B/op	       4 allocs/op
BenchmarkFlowWith2Param-8        	 2801313	       421.3 ns/op	     623 B/op	      12 allocs/op
BenchmarkFiberWith2Param-8       	  766850	      1394 ns/op	    2251 B/op	      26 allocs/op
BenchmarkGorillaWith2Param-8     	 1444689	       818.7 ns/op	    1182 B/op	      10 allocs/op
BenchmarkBunRouterWith2Param-8   	 5276590	       224.0 ns/op	     122 B/op	       4 allocs/op

BenchmarkKmuxWith5Param-8        	 2727201	       435.9 ns/op	     338 B/op	       9 allocs/op
BenchmarkKsmuxWith5Param-8       	 3522550	       337.6 ns/op	     252 B/op	       7 allocs/op
BenchmarkChiWith5Param-8         	 2125617	       564.8 ns/op	     638 B/op	       9 allocs/op
BenchmarkNetHTTPWith5Param-8     	  780470	      1469 ns/op	    1016 B/op	      17 allocs/op
BenchmarkGinWith5Param-8         	 3091296	       382.6 ns/op	     294 B/op	       7 allocs/op
BenchmarkEchoWith5Param-8        	 3134265	       379.4 ns/op	     261 B/op	       7 allocs/op
BenchmarkFlowWith5Param-8        	 1306460	       917.5 ns/op	    1062 B/op	      27 allocs/op
BenchmarkFiberWith5Param-8       	  747988	      1390 ns/op	    2211 B/op	      26 allocs/op
BenchmarkGorillaWith5Param-8     	  692682	      1558 ns/op	    1329 B/op	      13 allocs/op
BenchmarkBunRouterWith5Param-8   	 4723983	       253.5 ns/op	     152 B/op	       4 allocs/op
PASS
coverage: 0.0% of statements
```
