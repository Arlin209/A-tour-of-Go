package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Liftoff! 🚀")
}
