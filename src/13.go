package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	if x%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
