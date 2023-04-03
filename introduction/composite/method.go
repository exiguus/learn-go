package composite

import (
	"fmt"
	"math"
)

func (point Point) DistanceFromZero() float64 {
	return math.Sqrt(math.Pow(float64(point.X), 2) + math.Pow(float64(point.Y), 2))
}

func (point Point) String() string {
	return fmt.Sprintf("(%d | %d)", point.X, point.Y)
}
