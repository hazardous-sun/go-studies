package main

import "fmt"

func updateName(n *string) {
	*n = "wedge"
}

func main() {
	name := "tifa"

	updateName(&name)

	fmt.Println("memory address of name is:", &name)
	fmt.Println(name)
}
