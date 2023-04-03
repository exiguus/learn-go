package calculator

import "errors"

func Add(left, right int) int {
	return left + right
}

func Divide(left, right int) (int, int, error) {
	if right == 0 {
		return 0, 0, errors.New("Cannot divide by zero")
	}

	return left / right, left % right, nil
}
