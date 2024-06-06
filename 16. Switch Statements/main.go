package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	mybill := createBill()
	promptOptions(&mybill)

	fmt.Println("mybill.tip =", mybill.tip)

	fmt.Println(mybill.format())
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)

	fmt.Println("Created the bill -", b)

	return b
}

func promptOptions(b *bill) {
	reader := bufio.NewReader(os.Stdin)

	for true {
		opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

		switch opt {
		case "a":
			itemName, _ := getInput("Item name: ", reader)
			itemPrice := getPrice(reader)
			b.addItem(itemName, itemPrice)
		case "s":
			return
		case "t":
			tip := getPrice(reader)
			b.updateTip(tip)
		default:
			fmt.Println("Invalid option")
			promptOptions(b)
		}
	}
}

func getPrice(r *bufio.Reader) float64 {
	input, _ := getInput("Amount ($): ", r)
	price, _ := strconv.ParseFloat(input, 64)
	return price
}
