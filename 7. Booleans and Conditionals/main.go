package main

import "fmt"

func main() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("pikachu")
	} else if age < 40 {
		fmt.Println("agumon")
	} else {
		fmt.Println("tarkov")
	}

	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos:", index)
			continue
		}

		if index > 2 {
			fmt.Println("Breaking at pos:", index)
			break
		}

		fmt.Printf("the value at pos %v is %v\n", index, value)
	}
}
