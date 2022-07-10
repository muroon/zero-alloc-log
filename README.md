# zero-alloc-log

The simple zero allocation Logger in golang.
This is a sample for studying why the logger like [zerolog](https://github.com/rs/zerolog) or [zap](https://github.com/uber-go/zap) can be zero allocation.

Thease results are under Go 1.17.

## normallog package

[normallog](https://github.com/muroon/zero-alloc-log/tree/master/normallog) package is a simple log package.

```go
go test -cpu=1,2,4 -benchmem -benchtime=5s -bench . ./normallog/...
pkg: github.com/muroon/zero-alloc-log/normallog
BenchmarkEvent                  	221338058	        26.88 ns/op	       8 B/op	       1 allocs/op
BenchmarkEvent-2                	234270459	        23.69 ns/op	       8 B/op	       1 allocs/op
BenchmarkEvent-4                	243652970	        24.75 ns/op	       8 B/op	       1 allocs/op
BenchmarkEventWithDone          	168686601	        35.62 ns/op	       8 B/op	       1 allocs/op
BenchmarkEventWithDone-2        	235201628	        30.45 ns/op	       8 B/op	       1 allocs/op
BenchmarkEventWithDone-4        	226523143	        25.75 ns/op	       8 B/op	       1 allocs/op
BenchmarkEventWithInterface     	183410749	        33.43 ns/op	       8 B/op	       1 allocs/op
BenchmarkEventWithInterface-2   	226016618	        25.09 ns/op	       8 B/op	       1 allocs/op
BenchmarkEventWithInterface-4   	235461999	        25.50 ns/op	       8 B/op	       1 allocs/op
```

## poollog package

[poollog](https://github.com/muroon/zero-alloc-log/tree/master/poollog) package is a simple zero allocation log using sync.Pool. 

```go
go test -cpu=1,2,4 -benchmem -benchtime=5s -bench . ./poollog/...
pkg: github.com/muroon/zero-alloc-log/poollog
BenchmarkEvent                  	194794797	        33.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkEvent-2                	136148931	        44.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkEvent-4                	243642721	        24.51 ns/op	       0 B/op	       0 allocs/op
BenchmarkEventWithDone          	155714641	        36.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkEventWithDone-2        	204912210	        29.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkEventWithDone-4        	220397934	        28.56 ns/op	       0 B/op	       0 allocs/op
BenchmarkEventWithInterface     	178993027	        32.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkEventWithInterface-2   	221753439	        28.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkEventWithInterface-4   	194662105	        33.00 ns/op	       0 B/op	       0 allocs/op
```