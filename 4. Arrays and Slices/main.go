package main

import "fmt"

func main() {
	var ages [3]int = [3]int{1, 2, 3}
	var numbers = [4]int{10, 20, 30, 40}

	fmt.Println(ages, numbers)
	fmt.Println(len(ages), len(numbers))

	// slices
	var scores = []int{100, 200, 5}
	scores = append(scores, 100) // apparently, you need to assign the value of the slice again when appending

	names := [4]string{"John", "Paul", "George", "Ringo"}
	names[1] = "luigi"
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "koopa")

	fmt.Println(rangeOne)
}
