package main

import (
	"fmt"
	"log"
)

func main() {
	var input int
	fmt.Print("Enter a number:")
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatalln(err)
	}

	gen := generate(input)
	squared := square(gen)
	printInts(squared)
}
