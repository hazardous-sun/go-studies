package main

import "fmt"

func main() {
	// Print
	fmt.Print("hello, ")
	fmt.Print("world!\n")
	fmt.Print("hello again\n")

	// Println
	fmt.Println("hello dudes!")
	fmt.Println("goodbye dudes!")

	age := 35
	name := "shrek"
	fmt.Println("my age is", age)
	fmt.Println("name is", name)

	fmt.Printf("my name is %v and I am %v\n", name, age)
	fmt.Printf("my name is %q and I am %q\n", name, age)
	fmt.Printf("my name is %T and I am %T\n", name, age)
	fmt.Printf("you scored %0.2f points!\n", 255.5555)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and I am %v\n", name, age)
	fmt.Println(str)
}
