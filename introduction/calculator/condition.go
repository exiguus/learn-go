package calculator

import (
	"errors"
	"math"
)

func Abs(value int) int {
	if value >= 0 {
		return value
	}
	return -value
}

func SquareRoot(value int) (int, error) {
	if value < 0 {
		return 0, errors.New("value must not be negative")
	}
	return int(math.Sqrt(float64(value))), nil
}

func IsSquareRoot(value int) bool {
	if sqrt := math.Sqrt(float64(value)); sqrt != float64(int(sqrt)) {
		return false
	} else {
		return true
	}
}

func RunOperation(operation string, left, right int) int {
	var result int
	switch operation {
	case "add":
		result = Add(left, right)
	case "subtract":
		result = Subtract(left, right)
	}
	return result
}
