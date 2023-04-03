package calculator_test

import (
	"testing"

	"github.com/exiguus/gotesting/calculator"
)

func BenchmarkIsPrime(b *testing.B) {
	// setup
	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calculator.IsPrime(1000000007)
	}
}

func BenchmarkIsPrimeSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculator.IsPrime(18)
	}
}
