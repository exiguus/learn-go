# Go testing

## Run

```bash
go test ./...
# ?       github.com/exigus/gotesting     [no test files]
# ok      github.com/exigus/gotesting/calculator  0.004s
# ok      github.com/exigus/gotesting/network     (cached)
```

## Run with coverage

```bash
go test ./... -cover

# ?       github.com/exigus/gotesting     [no test files]
# ok      github.com/exigus/gotesting/calculator  0.004s  coverage: 100.0% of statements
# ok      github.com/exigus/gotesting/network     0.003s  coverage: 75.0% of statements
```

## Run with coverage and html output

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Run with coverage and html output and open in browser and delete coverage.out

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
xdg-open coverage.html
rm coverage.out
```

## Benchmark tests

```bash
go test -bench=. ./...
# goos: linux
# goarch: amd64
# pkg: github.com/exigus/gotesting/calculator
# cpu: Intel(R) Core(TM) i5-7300U CPU @ 2.60GHz
# BenchmarkIsPrime-4                  3717            293716 ns/op
# BenchmarkIsPrimeSmall-4         125128405                9.648 ns/op
# PASS
# ok      github.com/exigus/gotesting/calculator  3.308s
# PASS
# ok      github.com/exigus/gotesting/network     0.007s
```
