package main

import "fmt"

func printInts(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}
