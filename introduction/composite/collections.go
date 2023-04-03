package composite

import "fmt"

func DemoCollections() {
	primesArray := [5]int{2, 3, 5, 7, 11}
	primesSlice := make([]int, 5, 7)

	fmt.Println("DemoCollections")
	fmt.Println("array", primesArray)
	fmt.Println("array len", len(primesArray))
	fmt.Println("slice", primesSlice)
	fmt.Println("slice len", len(primesSlice))

	primesSlice[0] = 2
	primesSlice[1] = 3
	primesSlice[2] = 5
	primesSlice[3] = 7
	primesSlice = append(primesSlice, 11)
	fmt.Println("slice appended", primesSlice)
	fmt.Println("slice len", len(primesSlice))
	fmt.Println("slice cap", cap(primesSlice))

	for index, value := range primesSlice {
		fmt.Println("slice index, value", index, value)
	}

	primesMap := make(map[string]int)
	primesMap["two"] = 2
	primesMap["three"] = 3
	primesMap["five"] = 5
	primesMap["seven"] = 7
	primesMap["eleven"] = 11

	fmt.Println("map", primesMap)

	for key, value := range primesMap {
		fmt.Println("map key, value", key, value)
	}

	points := map[string]Point{
		"A": New(1, 2),
		"B": New(3, 4),
	}

	fmt.Println("points", points)

	fmt.Println("points A", points["A"])
	fmt.Println("points B", points["B"])
	fmt.Println("points C", points["C"])
	somePoint, ok := points["C"]
	fmt.Println("ask points ok for C", somePoint, ok)
	delete(points, "B")
	askForPointB, ok := points["B"]
	fmt.Println("delete point B in points and ask for B", askForPointB, ok)
	fmt.Println("points", points)

}
