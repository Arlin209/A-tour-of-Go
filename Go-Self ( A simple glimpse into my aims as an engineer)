package main

import "fmt"

// Function to display a motivational message
func motivate(name string, goal string) {
    fmt.Printf("Keep going %s! You're on your way to mastering %s.\n", name, goal)
}

func main() {
    // Personalized details
    name := "Arlin Wilson"
    college := "Christ College of Engineering, Irinjalakuda"
    field := "Electronics and Communication Engineering"
    hobby := "learning Go and building Embedded projects"

    // Print introduction
    fmt.Println("=== Personalized Go Program ===")
    fmt.Printf("Hello, my name is %s.\n", name)
    fmt.Printf("I study %s at %s.\n", field, college)
    fmt.Printf("One of my current interests is %s.\n", hobby)

    // Loop example
    fmt.Println("\nHere are my top goals:")
    goals := []string{
        "Complete EMG-based assistive device project",
        "Master C, Python, and Go programming",
        "Design a PCB",
        "Make engineering creative and fun",
    }

    // Loop through goals
    for i, g := range goals {
        fmt.Printf("%d. %s\n", i+1, g)
    }

    // Function call
    motivate(name, "Go programming")
}
