package main

import "fmt"

func main() {
	x := 42
	y := 13

	if x > y {
		fmt.Println("The value of x is greater than y")
	} else if x == y {
		fmt.Println("The value of x and y are equal")
	} else {
		fmt.Println("The value of y is greater than x")
	}
}
