package main

import (
	"fmt"
	"sort"
)

func main() {

	// strings
	//greetings := "farewell, my friend... will be together soon"
	//fmt.Println(strings.Contains(greetings, "farewell"))
	//fmt.Println(strings.Contains(greetings, "farewell!"))
	//fmt.Println(strings.ReplaceAll(greetings, "farewell", "hello"))
	//fmt.Println(strings.ToUpper(greetings))
	//fmt.Println(strings.Index(greetings, "will"))
	//fmt.Println(strings.Split(greetings, " "))

	// sort
	ages := []int{45, 20, 35, 30, 75, 60, 20, 50}
	sort.Ints(ages) // sorts a slice of integers
	fmt.Println(ages)

	index := sort.SearchInts(ages, 60)
	fmt.Println(index)

	index2 := sort.SearchInts(ages, 1) // The return value is the index to insert x if x is not present
	fmt.Printf("index2 value: %v\ttype: %T\n", index2, index2)

	names := []string{"John", "Paul", "George", "Ringo"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "George"))
}
