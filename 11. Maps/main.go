package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":                      4.99,
		"pie":                       7.99,
		"salad":                     6.99,
		"the blood of your enemies": 0.00,
	}

	fmt.Println(menu)
	fmt.Println(menu["salad"])

	// looping maps
	for k, v := range menu {
		fmt.Printf("%v: {%v}\n", k, v)
	}

	// ints as key type
	phonebook := map[int]string{
		123456789: "mario",
		987654321: "luigi",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[987654321])

	phonebook[123456789] = "danger zone"
	fmt.Println(phonebook[123456789])
}
