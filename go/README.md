# go-for-statistics
Go regression calculation and performance comparisons for the Anscombe Quartet datasets

## How to Use 
Open the repo, run the tests and benchmarks using this command:
```bash
go -C go test ./... -v -bench=. -benchmem -count=1
```


## Benchmark Results

### Go
```
goos: darwin
goarch: arm64
pkg: github.com/gummybeardream/go-for-statistics/anscombe
cpu: Apple M1
BenchmarkCalculateRegression/Dataset_I-8         	12504889	        95.44 ns/op	     448 B/op	       3 allocs/op
BenchmarkCalculateRegression/Dataset_II-8        	13054521	        96.36 ns/op	     448 B/op	       3 allocs/op
BenchmarkCalculateRegression/Dataset_III-8       	11931700	        97.58 ns/op	     448 B/op	       3 allocs/op
BenchmarkCalculateRegression/Dataset_IV-8        	12211124	        98.61 ns/op	     448 B/op	       3 allocs/op
PASS
ok  	github.com/gummybeardream/go-for-statistics/anscombe	5.235s
```

### Python
```
```

## AI Disclosure 
I used ChatGPT to teach me about testing and benchmarking in Go. I used it to write benchmarks in Python. 