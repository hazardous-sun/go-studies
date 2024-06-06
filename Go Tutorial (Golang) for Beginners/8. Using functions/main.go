package main

import (
	"fmt"
	"strings"
)

func addValues(x int, y int) int {
	return x + y
}

func messString(x string) string {
	x = strings.Replace(x, " ", "!", -1)
	x = strings.ToUpper(x)
	return x
}

func monoid(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func category_of_endofunctors(s string) {
	fmt.Println("a monad is just a monoid in the category of endofunctors, right", s+"?")
}

func main() {
	x := 1
	y := 2
	z := addValues(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, z)

	a := "Lorem Ipsum is definitely something that exists!"
	b := messString(a)

	fmt.Printf("original string = %v\tmessed string = %v\n", a, b)

	names := []string{"dory", "nemo", "stallings"}
	monoid(names, category_of_endofunctors)
}
