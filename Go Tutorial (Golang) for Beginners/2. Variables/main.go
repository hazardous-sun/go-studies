package main

import "fmt"

func main() {
	var nameOne string = "harry potato"
	var nameTwo = "gandalf the white"
	var nameThree string // does not hold invalid data, since it is initialized with ""

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "saruman"
	nameThree = "roberto carlos"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi" // only use ":" the first time
	fmt.Println(nameFour)

	// integers
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	var numOne int8 = 127
	var numTwo int16 = 32767
	var numThree int32 = 2147483647
	var numFour int64 = 9223372036854775807
	var numFive uint = 9223372036854775807 * 2

	fmt.Println(numOne, numTwo, numThree, numFour, numFive)

	var scoreOne float32 = 24.7756
	var scoreTwo float64 = 8323873412312939.1923812821983123899812398192083
	var scoreThree = 1293.22 // float64 by default

	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
