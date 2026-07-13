package main

import "fmt"

func main() {
    a, b := 10.0, 5.0
    op := "mul"

    var result float64
    switch op {
    case "add":
        result = a + b
    case "sub":
        result = a - b
    case "mul":
        result = a * b
    case "div":
        if b == 0 {
            fmt.Println("Error: Division by zero")
            return
        }
        result = a / b
    default:
        fmt.Println("Invalid operation")
        return
    }

    fmt.Printf("Result: %v\n", result)
}
