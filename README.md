# go-routers-benchmark

```shell
go test -benchmem -run=^$ -coverprofile=$(mktemp -t go-routers-benchmark) -bench . github.com/smhmayboudi/go-routers-benchmark
```

```shell
goos: darwin
goarch: arm64
pkg: github.com/smhmayboudi/go-routers-benchmark
BenchmarkBunRouter-8             	30206701	        39.65 ns/op	      17 B/op	       0 allocs/op
BenchmarkChi-8                   	 7318312	       147.9 ns/op	     359 B/op	       3 allocs/op
BenchmarkEcho-8                  	12788620	        93.56 ns/op	      18 B/op	       1 allocs/op
BenchmarkFiber-8                 	  772353	      1310 ns/op	    2110 B/op	      20 allocs/op
BenchmarkFlow-8                  	 5171496	       215.8 ns/op	     413 B/op	       2 allocs/op
BenchmarkGin-8                   	16132048	        72.82 ns/op	      64 B/op	       1 allocs/op
BenchmarkGorilla-8               	 2727082	       420.5 ns/op	     796 B/op	       7 allocs/op
BenchmarkKmux-8                  	 9424675	       126.8 ns/op	      86 B/op	       3 allocs/op
BenchmarkKsmux-8                 	25884242	        44.85 ns/op	      18 B/op	       1 allocs/op
BenchmarkNetHTTP-8               	 3877314	       310.0 ns/op	      22 B/op	       1 allocs/op

BenchmarkBunRouterWith1Param-8   	 5561899	       213.1 ns/op	     112 B/op	       4 allocs/op
BenchmarkChiWith1Param-8         	 4816305	       239.9 ns/op	     411 B/op	       5 allocs/op
BenchmarkEchoWith1Param-8        	 8170096	       144.8 ns/op	      80 B/op	       3 allocs/op
BenchmarkFiberWith1Param-8       	 1063873	      1142 ns/op	    2075 B/op	      22 allocs/op
BenchmarkFlowWith1Param-8        	 4491123	       261.2 ns/op	     477 B/op	       7 allocs/op
BenchmarkGinWith1Param-8         	 7725709	       153.6 ns/op	     114 B/op	       3 allocs/op
BenchmarkGorillaWith1Param-8     	 2012173	       582.3 ns/op	    1137 B/op	       9 allocs/op
BenchmarkKmuxWith1Param-8        	 5656089	       208.1 ns/op	     103 B/op	       5 allocs/op
BenchmarkKsmuxWith1Param-8       	 9763737	       121.8 ns/op	      75 B/op	       3 allocs/op
BenchmarkNetHTTPWith1Param-8     	 1795338	       667.6 ns/op	     440 B/op	      11 allocs/op

BenchmarkBunRouterWith2Param-8   	 5248740	       226.5 ns/op	     123 B/op	       4 allocs/op
BenchmarkChiWith2Param-8         	 3792876	       307.3 ns/op	     486 B/op	       6 allocs/op
BenchmarkEchoWith2Param-8        	 5916792	       201.7 ns/op	     125 B/op	       4 allocs/op
BenchmarkFiberWith2Param-8       	  864771	      1384 ns/op	    2231 B/op	      26 allocs/op
BenchmarkFlowWith2Param-8        	 2822461	       421.8 ns/op	     623 B/op	      12 allocs/op
BenchmarkGinWith2Param-8         	 5655002	       214.0 ns/op	     159 B/op	       4 allocs/op
BenchmarkGorillaWith2Param-8     	 1466680	       806.7 ns/op	    1182 B/op	      10 allocs/op
BenchmarkKmuxWith2Param-8        	 4507852	       263.0 ns/op	     179 B/op	       6 allocs/op
BenchmarkKsmuxWith2Param-8       	 6838110	       172.7 ns/op	     119 B/op	       4 allocs/op
BenchmarkNetHTTPWith2Param-8     	 1355002	       885.0 ns/op	     536 B/op	      13 allocs/op

BenchmarkBunRouterWith5Param-8   	 4738866	       252.2 ns/op	     152 B/op	       4 allocs/op
BenchmarkChiWith5Param-8         	 2118760	       559.2 ns/op	     638 B/op	       9 allocs/op
BenchmarkEchoWith5Param-8        	 3110529	       383.4 ns/op	     262 B/op	       7 allocs/op
BenchmarkFiberWith5Param-8       	  720205	      1394 ns/op	    2216 B/op	      26 allocs/op
BenchmarkFlowWith5Param-8        	 1297449	       932.9 ns/op	    1063 B/op	      27 allocs/op
BenchmarkGinWith5Param-8         	 3030270	       401.9 ns/op	     296 B/op	       7 allocs/op
BenchmarkGorillaWith5Param-8     	  656086	      1591 ns/op	    1334 B/op	      13 allocs/op
BenchmarkKmuxWith5Param-8        	 2719119	       442.5 ns/op	     338 B/op	       9 allocs/op
BenchmarkKsmuxWith5Param-8       	 3469557	       342.0 ns/op	     253 B/op	       7 allocs/op
BenchmarkNetHTTPWith5Param-8     	  801680	      1471 ns/op	    1016 B/op	      17 allocs/op
PASS
coverage: 0.0% of statements
```
