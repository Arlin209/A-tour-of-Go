package main

import "fmt"

// Function to multiply two numbers
func multiply(a int, b int) int {
    return a * b
}

func main() {
    fmt.Println("Simple Go Program")

    // Variables
    x := 6
    y := 7

    // Function call
    result := multiply(x, y)

    // Output
    fmt.Printf("The product of %d and %d is %d\n", x, y, result)
}
