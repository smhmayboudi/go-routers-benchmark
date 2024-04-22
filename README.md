# go-routers-benchmark

```shell
go test -benchmem -run=^$ -coverprofile=$(mktemp -t go-routers-benchmark) -bench . github.com/smhmayboudi/go-routers-benchmark
```

```shell
goos: darwin
goarch: arm64
pkg: github.com/smhmayboudi/go-routers-benchmark
BenchmarkBunRouter-8              	30125264	        40.32 ns/op	      17 B/op	       0 allocs/op
BenchmarkChi-8                    	 7263196	       150.6 ns/op	     359 B/op	       3 allocs/op
BenchmarkEcho-8                   	12569036	        94.87 ns/op	      18 B/op	       1 allocs/op
BenchmarkFiber-8                  	  793477	      1324 ns/op	    2107 B/op	      20 allocs/op
BenchmarkFlow-8                   	 5189262	       217.4 ns/op	     412 B/op	       2 allocs/op
BenchmarkGin-8                    	16072962	        74.02 ns/op	      64 B/op	       1 allocs/op
BenchmarkGorilla-8                	 2665440	       424.0 ns/op	     796 B/op	       7 allocs/op
BenchmarkHttpRouter-8             	31083820	        40.38 ns/op	      17 B/op	       0 allocs/op
BenchmarkKmux-8                   	 9166278	       128.3 ns/op	      86 B/op	       3 allocs/op
BenchmarkKsmux-8                  	26108168	        45.19 ns/op	      18 B/op	       1 allocs/op
BenchmarkNetHTTP-8                	 3844714	       313.3 ns/op	      22 B/op	       1 allocs/op

BenchmarkBunRouterWith1Param-8    	13255353	        91.06 ns/op	      56 B/op	       1 allocs/op
BenchmarkChiWith1Param-8          	 4729028	       242.1 ns/op	     412 B/op	       5 allocs/op
BenchmarkEchoWith1Param-8         	 8226380	       144.9 ns/op	      80 B/op	       3 allocs/op
BenchmarkFiberWith1Param-8        	 1036297	      1157 ns/op	    2078 B/op	      22 allocs/op
BenchmarkFlowWith1Param-8         	 4418326	       261.8 ns/op	     478 B/op	       7 allocs/op
BenchmarkGinWith1Param-8          	 7660002	       155.0 ns/op	     115 B/op	       3 allocs/op
BenchmarkGorillaWith1Param-8      	 1990021	       590.1 ns/op	    1138 B/op	       9 allocs/op
BenchmarkHttpRouterWith1Param-8   	12244044	        97.46 ns/op	      91 B/op	       2 allocs/op
BenchmarkKmuxWith1Param-8         	 5708355	       210.8 ns/op	     103 B/op	       5 allocs/op
BenchmarkKsmuxWith1Param-8        	 9587133	       121.5 ns/op	      76 B/op	       3 allocs/op
BenchmarkNetHTTPWith1Param-8      	 1771501	       674.1 ns/op	     440 B/op	      11 allocs/op

BenchmarkBunRouterWith2Param-8    	 7109019	       164.4 ns/op	      69 B/op	       2 allocs/op
BenchmarkChiWith2Param-8          	 3773150	       310.9 ns/op	     487 B/op	       6 allocs/op
BenchmarkEchoWith2Param-8         	 5897835	       200.3 ns/op	     125 B/op	       4 allocs/op
BenchmarkFiberWith2Param-8        	  721528	      1492 ns/op	    2216 B/op	      26 allocs/op
BenchmarkFlowWith2Param-8         	 2790025	       424.5 ns/op	     624 B/op	      12 allocs/op
BenchmarkGinWith2Param-8          	 5651860	       220.3 ns/op	     159 B/op	       4 allocs/op
BenchmarkGorillaWith2Param-8      	 1454626	       815.1 ns/op	    1182 B/op	      10 allocs/op
BenchmarkHttpRouterWith2Param-8   	 7929288	       150.4 ns/op	     163 B/op	       3 allocs/op
BenchmarkKmuxWith2Param-8         	 4539235	       265.1 ns/op	     179 B/op	       6 allocs/op
BenchmarkKsmuxWith2Param-8        	 6785292	       173.4 ns/op	     119 B/op	       4 allocs/op
BenchmarkNetHTTPWith2Param-8      	 1339316	       896.0 ns/op	     536 B/op	      13 allocs/op

BenchmarkBunRouterWith5Param-8    	 2663917	       447.0 ns/op	     180 B/op	       5 allocs/op
BenchmarkChiWith5Param-8          	 2093282	       563.0 ns/op	     640 B/op	       9 allocs/op
BenchmarkEchoWith5Param-8         	 3106786	       382.0 ns/op	     262 B/op	       7 allocs/op
BenchmarkFiberWith5Param-8        	  717952	      1407 ns/op	    2217 B/op	      26 allocs/op
BenchmarkFlowWith5Param-8         	 1289635	       926.7 ns/op	    1064 B/op	      27 allocs/op
BenchmarkGinWith5Param-8          	 3088273	       386.6 ns/op	     294 B/op	       7 allocs/op
BenchmarkGorillaWith5Param-8      	  694713	      1569 ns/op	    1328 B/op	      13 allocs/op
BenchmarkHttpRouterWith5Param-8   	 3618318	       326.5 ns/op	     314 B/op	       6 allocs/op
BenchmarkKmuxWith5Param-8         	 2739228	       435.0 ns/op	     338 B/op	       9 allocs/op
BenchmarkKsmuxWith5Param-8        	 3504487	       338.9 ns/op	     252 B/op	       7 allocs/op
BenchmarkNetHTTPWith5Param-8      	  797052	      1486 ns/op	    1016 B/op	      17 allocs/op
PASS
coverage: 0.0% of statements
```
