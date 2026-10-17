# go-for-statistics
Go regression calculation and performance comparisons for the Anscombe Quartet datasets

## How to Use 
Open the repo, from the top level directory, run the tests and benchmarks using this command for Go:
```bash
go -C go test ./... -v -bench=. -benchmem -count=1
```
Use this command for Python: 
```python
python python/benchmark.py
```

If you want to run the Go code, use the command: 
```bash
go -C go run ./...
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
```Dataset I:
  Execution time: 58534.13 ns/op
  Peak traced memory: 7736 bytes
Dataset II:
  Execution time: 53505.56 ns/op
  Peak traced memory: 7736 bytes
Dataset III:
  Execution time: 53046.52 ns/op
  Peak traced memory: 7736 bytes
Dataset IV:
  Execution time: 52581.25 ns/op
  Peak traced memory: 7736 bytes
```

### Recommendation 

Comparing the benchmark results, Go executed the regression calculation at 95-99ns/op versus 53,000-59,000 ns/op for Python. Go's montanaflynn/stats library was faster than Python's Statsmodel when calculating the four regressions for the Anscombe Quartet. Both libraries were able to produce closely matching regression coefficients for all four datasets. 

Since measurements for Go show fast execution for the simple regression calculations, this would be valuable when an application needs to perform many repeated calculations such as processing incoming data for backend research and product development. More testing with large datasets and high volume workloads would yield more meaningful performance comparisons. 

While each implementation's memory measurements were consistent across all four datasets, Go's allocated bytes per operation and Python's peak traced memory are different measurements. 

For the data science team, I would recommend Go for faster computational performance and high volume simple calculations but Python would be more practical for more advanced statistical analysis. 

## AI Disclosure 
I used ChatGPT to teach me about testing and benchmarking in Go while writing it by hand. I used ChatGPT to guide me through writing benchmarks in Python for the provided script. 