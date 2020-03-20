package main

import (
	"fmt"
	"time"
)

func longop(duration int) {
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Printf("Took %d secs\n", duration)
}

func main() {

	fmt.Println("Starting first operation :")
	go longop(2)
	fmt.Println("Starting second operation :")
	longop(2)

	time.Sleep(3 * time.Second)
}
