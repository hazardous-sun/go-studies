package main

import "fmt"

func updateName1(x string) {
	x = "wedge"
}

func updateName2(x string) string {
	x = "wedge"
	return x
}

func updateMenu(x map[string]float64) {
	x["coffe"] = 2.99
}

func main() {
	// group A types (Non-Pointer Values) -> strings, ints, bools, floats, arrays, structs (passed by value)
	// weird decision to clone strings by default...
	name := "tifa"
	updateName1(name) // clones the value to the function, not passing it by reference
	fmt.Println(name)

	name = updateName2(name)
	fmt.Println(name)

	// group B types (Pointer Wrapper Values) -> slices, maps, functions (passed by reference)
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)
}
