package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) []string {
	split := strings.Split(strings.ToUpper(n), " ")
	var initials []string
	for _, name := range split {
		if len(name) > 0 {
			initials = append(initials, name[:1])
		}
	}
	return initials
}

func namesInitials(n []string) [][]string {
	var initials [][]string
	for _, name := range n {
		initials = append(initials, getInitials(name))
	}
	return initials
}

func twoWordsInitials(n string) (string, string) {
	names := strings.Split(strings.ToUpper(n), " ")
	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	names := []string{"Tifa Lockhart", "Gray Fullbuster", "Monkey D. Luffy", "Spongebob Squarepants", "Silvio Santos"}
	fmt.Println(namesInitials(names))

	name := "Roger Rabbit"
	fmt.Println(getInitials(name))

	name = "Roger "
	fmt.Println(getInitials(name))

	fn, sn := twoWordsInitials("Tifa Lockhart")
	fmt.Println(fn, sn)
}
