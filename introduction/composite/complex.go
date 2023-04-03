package composite

func Add(a, b int) int {
	aPtr := &a
	bPtr := &b

	sum := *aPtr + *bPtr
	return sum
}
