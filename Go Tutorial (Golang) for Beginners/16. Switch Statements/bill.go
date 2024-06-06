package main

import (
	"fmt"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	return bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
}

// format the bill
func (b *bill) format() string {
	formatedString := "Bill breakdown: \n"

	var total float64 = 0

	// list items
	for k, v := range b.items {
		formatedString += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	// tip
	formatedString += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)
	total += b.tip

	// total
	formatedString += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return formatedString
}

// update the tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
