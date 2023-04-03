package calculator_test

import (
	"testing"

	"github.com/exiguus/gotesting/calculator"
	"github.com/stretchr/testify/assert"
)

// func TestAdd(t *testing.T) {
// 	actual := calculator.Add(1, 2)
// 	expect := 3

// 	if actual != expect {
// 		t.Errorf("Wrong sum, expect %d but got %d", expect, actual)
// 	}
// }

// Table driven tests
func TestAdd(t *testing.T) {
	var tests = []struct {
		name                string
		left, right, expect int
	}{
		{"positive Integer", 1, 1, 2},
		{"negative Integer", -12, -6, -18},
		{"null Integer", 0, 0, 0},
		{"positive and negative Integer", 12, -6, 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, calculator.Add(test.left, test.right))
			// if actual := calculator.Add(test.x, test.y); actual != test.expect {
			// 	t.Errorf("Wrong sum, expect %d but got %d", test.expect, actual)
			// }
		})
	}
}

func TestDivide(t *testing.T) {
	var tests = []struct {
		name                string
		left, right         int
		quotient, remainder int
		withError           bool
	}{
		{"positive Integer", 1, 1, 1, 0, false},
		{"positive Integer remainder", 17, 3, 5, 2, false},
		{"negative Integer", -12, -6, 2, 0, false},
		{"positive and negative Integer", 12, -6, -2, 0, false},
		{"wrong", 12, 0, 0, 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			quotient, remainder, err := calculator.Divide(test.left, test.right)

			if test.withError {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, test.quotient, quotient)
			assert.Equal(t, test.remainder, remainder)
			assert.NoError(t, err)
			// if actual := calculator.Add(test.x, test.y); actual != test.expect {
			// 	t.Errorf("Wrong sum, expect %d but got %d", test.expect, actual)
			// }
		})
	}
}
