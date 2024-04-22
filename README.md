# go-routers-benchmark

```shell
go test -benchmem -run=^$ -coverprofile=$(mktemp -t go-routers-benchmark) -bench . github.com/smhmayboudi/go-routers-benchmark
```

```shell
goos: darwin
goarch: arm64
pkg: github.com/smhmayboudi/go-routers-benchmark

BenchmarkBunRouter-8              	29811811	        40.16 ns/op	      18 B/op	       0 allocs/op
BenchmarkChi-8                    	 7374037	       147.2 ns/op	     359 B/op	       3 allocs/op
BenchmarkEcho-8                   	12793647	        94.12 ns/op	      18 B/op	       1 allocs/op
BenchmarkFiber-8                  	  779364	      1310 ns/op	    2109 B/op	      20 allocs/op
BenchmarkFlow-8                   	 5222136	       216.2 ns/op	     412 B/op	       2 allocs/op
BenchmarkGin-8                    	15939076	        73.54 ns/op	      64 B/op	       1 allocs/op
BenchmarkGorilla-8                	 2711346	       418.8 ns/op	     796 B/op	       7 allocs/op
BenchmarkHttpRouter-8             	31583275	        38.88 ns/op	      16 B/op	       0 allocs/op
BenchmarkKmux-8                   	 9128037	       127.6 ns/op	      86 B/op	       3 allocs/op
BenchmarkKsmux-8                  	26112240	        44.81 ns/op	      18 B/op	       1 allocs/op
BenchmarkNetHTTP-8                	 3891087	       311.3 ns/op	      22 B/op	       1 allocs/op

BenchmarkBunRouterWith1Param-8    	13290403	        90.96 ns/op	      56 B/op	       1 allocs/op
BenchmarkChiWith1Param-8          	 4777564	       237.3 ns/op	     412 B/op	       5 allocs/op
BenchmarkEchoWith1Param-8         	 8262081	       143.7 ns/op	      80 B/op	       3 allocs/op
BenchmarkFiberWith1Param-8        	 1042105	      1149 ns/op	    2077 B/op	      22 allocs/op
BenchmarkFlowWith1Param-8         	 4493044	       259.9 ns/op	     477 B/op	       7 allocs/op
BenchmarkGinWith1Param-8          	 7657814	       153.8 ns/op	     115 B/op	       3 allocs/op
BenchmarkGorillaWith1Param-8      	 2011726	       582.7 ns/op	    1137 B/op	       9 allocs/op
BenchmarkHttpRouterWith1Param-8   	12596600	        95.34 ns/op	      90 B/op	       2 allocs/op
BenchmarkKmuxWith1Param-8         	 5750445	       208.2 ns/op	     103 B/op	       5 allocs/op
BenchmarkKsmuxWith1Param-8        	 9817094	       120.7 ns/op	      75 B/op	       3 allocs/op
BenchmarkNetHTTPWith1Param-8      	 1798729	       669.8 ns/op	     440 B/op	      11 allocs/op

BenchmarkBunRouterWith2Param-8    	 7257954	       161.8 ns/op	      68 B/op	       2 allocs/op
BenchmarkChiWith2Param-8          	 3841756	       304.9 ns/op	     485 B/op	       6 allocs/op
BenchmarkEchoWith2Param-8         	 5963230	       199.5 ns/op	     125 B/op	       4 allocs/op
BenchmarkFiberWith2Param-8        	  716545	      1399 ns/op	    2217 B/op	      26 allocs/op
BenchmarkFlowWith2Param-8         	 2812090	       419.4 ns/op	     623 B/op	      12 allocs/op
BenchmarkGinWith2Param-8          	 5622560	       215.2 ns/op	     159 B/op	       4 allocs/op
BenchmarkGorillaWith2Param-8      	 1467932	       812.6 ns/op	    1182 B/op	      10 allocs/op
BenchmarkHttpRouterWith2Param-8   	 7979810	       149.2 ns/op	     163 B/op	       3 allocs/op
BenchmarkKmuxWith2Param-8         	 4535203	       263.8 ns/op	     179 B/op	       6 allocs/op
BenchmarkKsmuxWith2Param-8        	 6765435	       173.9 ns/op	     119 B/op	       4 allocs/op
BenchmarkNetHTTPWith2Param-8      	 1354952	       889.4 ns/op	     536 B/op	      13 allocs/op

BenchmarkBunRouterWith5Param-8    	 2700884	       442.3 ns/op	     179 B/op	       5 allocs/op
BenchmarkChiWith5Param-8          	 2125104	       561.0 ns/op	     638 B/op	       9 allocs/op
BenchmarkEchoWith5Param-8         	 3123518	       381.0 ns/op	     261 B/op	       7 allocs/op
BenchmarkFiberWith5Param-8        	  710704	      1410 ns/op	    2218 B/op	      26 allocs/op
BenchmarkFlowWith5Param-8         	 1302312	       920.0 ns/op	    1063 B/op	      27 allocs/op
BenchmarkGinWith5Param-8          	 3070933	       393.8 ns/op	     295 B/op	       7 allocs/op
BenchmarkGorillaWith5Param-8      	  685748	      1560 ns/op	    1330 B/op	      13 allocs/op
BenchmarkHttpRouterWith5Param-8   	 3619641	       327.1 ns/op	     314 B/op	       6 allocs/op
BenchmarkKmuxWith5Param-8         	 2735426	       434.8 ns/op	     338 B/op	       9 allocs/op
BenchmarkKsmuxWith5Param-8        	 3515283	       339.8 ns/op	     252 B/op	       7 allocs/op
BenchmarkNetHTTPWith5Param-8      	  801085	      1466 ns/op	    1016 B/op	      17 allocs/op

PASS
coverage: 0.0% of statements
```
