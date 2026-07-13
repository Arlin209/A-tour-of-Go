package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    choices := []string{"rock", "paper", "scissors"}
    var player string

    fmt.Println("=== Rock-Paper-Scissors ===")
    fmt.Print("Enter rock, paper, or scissors: ")
    fmt.Scan(&player)

    computer := choices[rand.Intn(len(choices))]
    fmt.Printf("Computer chose: %s\n", computer)

    if player == computer {
        fmt.Println("It's a tie!")
    } else if (player == "rock" && computer == "scissors") ||
        (player == "paper" && computer == "rock") ||
        (player == "scissors" && computer == "paper") {
        fmt.Println("🎉 You win!")
    } else {
        fmt.Println("💻 Computer wins!")
    }
}
