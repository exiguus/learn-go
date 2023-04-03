package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/exiguus/gointroduction/calculator"
	"github.com/exiguus/gointroduction/composite"
)

func main() {
	a := rand.Intn(23)
	var b = 42

	fmt.Printf("Add %[1]d + %[2]d = %[3]d\n", a, b, calculator.Add(a, b))
	fmt.Printf("Subtract %[1]d - %[2]d = %[3]d\n", a, b, calculator.Subtract(a, b))
	fmt.Printf("SumUntil(%[1]d) = %[2]d\n", a*a*b, calculator.SumUntil(a*a*b))
	fmt.Printf("Sum(%[1]d, %[2]d) = %[3]d\n", 4, 123, calculator.Sum(4, 123))
	fmt.Printf("Abs(%[1]d) = %[2]d\n", a, calculator.Abs(a))
	fmt.Printf("isSquareRoot(%[1]d) = %[2]t\n", a, calculator.IsSquareRoot(a))
	fmt.Printf("RunOperation(%[1]q, %[2]d, %[3]d) = %[4]d\n", "add", a, b, calculator.RunOperation("add", a, b))
	fmt.Printf("RunOperation(%[1]q, %[2]d, %[3]d) = %[4]d\n", "subtract", a, b, calculator.RunOperation("subtract", a, b))
	fmt.Printf("Multiply(%[1]d, %[2]d) = %[3]d\n", a, b, calculator.Multiply(a, b))
	fmt.Printf("MultiplyFromAToB(%[1]d, %[2]d) = %[3]d\n", a, b, calculator.MultiplyFromAToB(a, b))
	fmt.Printf("SumFromAToB(%[1]d, %[2]d) = %[3]d\n", a, b, calculator.SumFromAToB(a, b))
	fmt.Printf("Add(%[1]d, %[2]d) = %[3]d\n", a, b, composite.Add(a, b))

	// Composite types and interfaces
	point := composite.New(a, b)

	composite.Print(point)
	composite.Print(point.X)
	composite.Print(point.Y)

	otherPoint := composite.Point{X: a * a, Y: b * b}
	otherPtr := &otherPoint

	fmt.Println(otherPtr, otherPtr.X, otherPtr.Y)

	// Collections
	// composite.DemoCollections()

	point = composite.New(3, 2)
	composite.Print(point.DistanceFromZero())
	composite.Print(false)
	composite.Print(1)
	composite.Print(point.String())
	composite.Print([]int{1, 2, 3})

	// Error handling
	result, err := calculator.SquareRoot(a * -1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("SquareRoot(%[1]d) = %[2]d\n", a, result)
	}

	// Goroutines

	queue := make(chan int, 100)
	go func(q chan int) {
		time.Sleep(3 * time.Second)
		fmt.Println("Hello from goroutine")
		q <- 23
		time.Sleep(3 * time.Second)
		q <- 42
		time.Sleep(3 * time.Second)
		q <- 69
		time.Sleep(3 * time.Second)
		q <- 666

	}(queue)

	valueFromQueue := <-queue
	fmt.Println(valueFromQueue)
}
