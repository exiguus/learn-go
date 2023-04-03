package calculator

func SumFromAToB(a, b int) int {
	return ProcessFromAToB(a, b, 0, func(x, y int) int {
		return Add(x, y)
	})
}

func MultiplyFromAToB(a, b int) int {
	return ProcessFromAToB(a, b, 1, func(x, y int) int {
		return Multiply(x, y)
	})
}

func ProcessFromAToB(a, b, initValue int, fn func(int, int) int) int {
	if a > b {
		return initValue
	}

	return fn(a, ProcessFromAToB(a+1, b, initValue, fn))
}
