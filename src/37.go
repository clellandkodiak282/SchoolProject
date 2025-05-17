package main

import "fmt"

func main() {
    var scores []float64 = [5]int{89.2, 75.1, 90.3, 68.1, 94.3}
    
    for i := range scores {
        fmt.Println(scores[i])
    }
}
