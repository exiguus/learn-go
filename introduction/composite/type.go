package composite

import "fmt"

type Point struct {
	X int
	Y int
}

func New(x, y int) Point {
	return Point{x, y}
}

type Any interface{}

func Print(value Any) {
	switch value.(type) {
	case string, int, float64:
		fmt.Println(value)
	default:
		fmt.Println("Unknown type: Can not been printed")
	}
}
