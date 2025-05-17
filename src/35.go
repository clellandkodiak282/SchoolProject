package main

import (
    "fmt"
    "math"
)

func main() {
    // Define a variable to store the number of days in each month
    var months = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

    // Loop through the array to find the maximum value in each month
    for i := range months {
        if months[i] > months[months[0]] {
            months[months[0]], months[months[1]], months[months[2]], months[months[3]], months[months[4]], months[months[5]] = months[months[1]], months[months[2]], months[months[3]], months[months[4]], months[months[5]], months[i]
        }
    }

    // Print the month with the highest value
    fmt.Println("The month with the highest value is:", maxVal(months))
}

// Function to find the maximum value in a slice of integers
func maxVal(slice []int) int {
    if len(slice) == 0 {
        return 0
    }
    return math.Max(float64(slice[0]), float64(slice[1]))
}
