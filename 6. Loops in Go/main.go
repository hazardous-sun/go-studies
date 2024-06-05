package main

import "fmt"

func main() {
	x := 0

	// apparently Go doesn't has a "while" structure, every loop is handled with "for"
	for x < 5 {
		fmt.Println("x = ", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("i = ", i)
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Println(index, value)
	}

	for _, value := range names {
		fmt.Println(value)
	}
}
