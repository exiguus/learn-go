package calculator_test

import (
	"testing"

	"github.com/exiguus/gotesting/calculator"
)

func TestIsPrime(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"negative", -1, false},
		{"zero", 0, false},
		{"one", 1, false},
		{"two", 2, true},
		{"three", 3, true},
		{"four", 4, false},
		{"five", 5, true},
		{"six", 6, false},
		{"seven", 7, true},
		{"eight", 8, false},
		{"nine", 9, false},
		{"ten", 10, false},
		{"eleven", 11, true},
		{"twelve", 12, false},
		{"thirteen", 13, true},
		{"fourteen", 14, false},
		{"fifteen", 15, false},
		{"sixteen", 16, false},
		{"seventeen", 17, true},
		{"eighteen", 18, false},
		{"nineteen", 19, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := calculator.IsPrime(test.n)
			if got != test.want {
				t.Errorf("Wrong sum, expect %t but got %t", test.want, got)
			}
		})
	}
}
