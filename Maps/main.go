package main

import "fmt"

func main() {
	// First way of declaring Map
	colors := map[string]string{
		"r": "Red",
		"g": "Green",
		"b": "Blue",
	}
	fmt.Println(colors)

	// Second way of declaring Map => NOT recomendad
	var days map[int]string
	// days[0] = "Monday" <-- This will cause Panic at runtime due to nil initalization
	// Map types are reference types, like pointers or slices, and so the value of above
	// is nil; it doesn't point to an initialized map. A nil map behaves like an
	// empty map when reading, but attempts to write to a nil map will cause a runtime
	// panic; don't do that. To initialize a map, use the built in make function
	fmt.Println(days)

	// Third way of declaring Map
	months := make(map[int]string)
	months[1] = "January"
	months[2] = "Febuary"
	months[3] = "March"
	
	iterMap(months)

	fmt.Println("Before delete", months)
	delete(months, 1)
	fmt.Println("After delete", months)
}

func iterMap(c map[int]string) {
	fmt.Println("Num -> Name")
	for idx, name := range c {
		fmt.Println(idx, " -> ", name)
	}
}
