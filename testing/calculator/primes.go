package calculator

import "math"

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	//
	//
	// unoptimized

	// goos: linux
	// goarch: amd64
	// pkg: github.com/exigus/gotesting/calculator
	// cpu: Intel(R) Core(TM) i5-7300U CPU @ 2.60GHz
	// BenchmarkIsPrime-4                     1        8365674085 ns/op
	// BenchmarkIsPrimeSmall-4         138018446                8.656 ns/op
	// PASS
	// ok      github.com/exigus/gotesting/calculator  10.451s
	// PASS
	// ok      github.com/exigus/gotesting/network     0.004s

	// for i := 2; i < n; i++ {
	// 	if n%i == 0 {
	// 		return false
	// 	}
	// }

	//
	//
	// optimized

	// goos: linux
	// goarch: amd64
	// pkg: github.com/exigus/gotesting/calculator
	// cpu: Intel(R) Core(TM) i5-7300U CPU @ 2.60GHz
	// BenchmarkIsPrime-4                  3542            323522 ns/op
	// BenchmarkIsPrimeSmall-4         100000000               10.51 ns/op
	// PASS
	// ok      github.com/exigus/gotesting/calculator  3.082s
	// PASS
	// ok      github.com/exigus/gotesting/network     0.005s
	if n == 2 {
		return true
	}

	for i := 2; i <= int(math.Ceil(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
