package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	for i, v := range args {
		fmt.Printf("i = %d\tv = %s\n", i, v)
	}
}
